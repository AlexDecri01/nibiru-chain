package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sort"

	"github.com/NibiruChain/nibiru/x/oracle/types"
)

// Tally calculates the median and returns it. Sets the set of voters to be rewarded, i.e. voted within
// a reasonable spread from the weighted median to the store
func Tally(
	ballot types.ExchangeRateBallot,
	rewardBand sdk.Dec,
	validatorPerformanceMap map[string]types.ValidatorPerformance,
) sdk.Dec {
	sort.Sort(ballot)

	weightedMedian := ballot.WeightedMedianWithAssertion()
	standardDeviation := ballot.StandardDeviation(weightedMedian)
	rewardSpread := weightedMedian.Mul(rewardBand.QuoInt64(2))

	if standardDeviation.GT(rewardSpread) {
		rewardSpread = standardDeviation
	}

	for _, vote := range ballot {
		// Filter ballot winners & abstain voters
		if (vote.ExchangeRate.GTE(weightedMedian.Sub(rewardSpread)) &&
			vote.ExchangeRate.LTE(weightedMedian.Add(rewardSpread))) ||
			!vote.ExchangeRate.IsPositive() {

			key := vote.Voter.String()
			// TODO: WTF is this? Mutating a map in Tally without no reason?
			claim := validatorPerformanceMap[key]
			claim.Weight += vote.Power
			claim.WinCount++
			validatorPerformanceMap[key] = claim
		}
	}

	return weightedMedian
}

// ballot for the asset is passing the threshold amount of voting power
func ballotIsPassing(ballot types.ExchangeRateBallot, thresholdVotes sdk.Int) bool {
	ballotPower := sdk.NewInt(ballot.Power())
	return !ballotPower.IsZero() && ballotPower.GTE(thresholdVotes)
}

// RemoveInvalidBallots removes the ballots which have not reached the vote threshold
// or which are not part of the vote targets anymore: example when params change during a vote period
// but some votes were already made.
func (k Keeper) RemoveInvalidBallots(ctx sdk.Context, voteTargets map[string]struct{}, voteMap map[string]types.ExchangeRateBallot) {
	totalBondedPower := sdk.TokensToConsensusPower(k.StakingKeeper.TotalBondedTokens(ctx), k.StakingKeeper.PowerReduction(ctx))
	voteThreshold := k.VoteThreshold(ctx)
	thresholdVotes := voteThreshold.MulInt64(totalBondedPower).RoundInt()

	for pair, ballot := range voteMap {
		// If pair is not in the voteTargets, or the ballot for it has failed, then skip
		// and remove it from voteMap for iteration efficiency
		if _, exists := voteTargets[pair]; !exists {
			delete(voteMap, pair)
			continue
		}

		// If the ballot is not passed, remove it from the voteTargets array
		// to prevent slashing validators who did valid vote.
		if !ballotIsPassing(ballot, thresholdVotes) {
			delete(voteTargets, pair)
			delete(voteMap, pair)
			continue
		}
	}
}
