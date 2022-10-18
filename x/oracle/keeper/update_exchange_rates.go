package keeper

import (
	"sort"

	"github.com/NibiruChain/nibiru/collections"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/NibiruChain/nibiru/x/oracle/types"
)

func (k Keeper) UpdateExchangeRates(ctx sdk.Context) {
	k.Logger(ctx).Info("processing validator price votes")

	pairsMap := k.getPairsMap(ctx)

	k.resetExchangeRates(ctx)

	validatorPerformanceMap := k.getValidatorPerformanceMap(ctx)

	pairBallotMap := k.OrganizeBallotByPair(ctx, validatorPerformanceMap)
	k.RemoveInvalidBallots(ctx, pairsMap, pairBallotMap)

	// Iterate through ballots and update exchange rates; drop if not enough votes have been achieved.
	params := k.GetParams(ctx)
	for pair, ballot := range pairBallotMap {
		sort.Sort(ballot)

		// Get weighted median of cross exchange rates
		exchangeRate := Tally(ctx, ballot, params.RewardBand, validatorPerformanceMap)

		// Set the exchange rate, emit ABCI event
		k.ExchangeRates.Insert(ctx, pair, exchangeRate)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.EventTypeExchangeRateUpdate,
				sdk.NewAttribute(types.AttributeKeyPair, pair),
				sdk.NewAttribute(types.AttributeKeyExchangeRate, exchangeRate.String()),
			),
		)
	}

	//---------------------------
	// Do miss counting & slashing
	voteTargetsLen := len(pairsMap)
	for _, claim := range validatorPerformanceMap {
		// Skip abstain & valid voters
		if int(claim.WinCount) == voteTargetsLen {
			continue
		}

		// Increase miss counter
		k.MissCounters.Insert(ctx, claim.ValAddress, k.MissCounters.GetOr(ctx, claim.ValAddress, 0)+1)
		k.Logger(ctx).Info("vote miss", "validator", claim.ValAddress.String())
	}

	// Distribute rewards to ballot winners
	k.RewardBallotWinners(ctx, pairsMap, validatorPerformanceMap)

	// Clear the ballot
	k.ClearBallots(ctx, params.VotePeriod)

	// Update vote targets
	k.ApplyWhitelist(ctx, params.Whitelist, pairsMap)
}

// getPairsMap returns a map containing all the pairs as the key.
func (k Keeper) getPairsMap(ctx sdk.Context) map[string]struct{} {
	pairsMap := make(map[string]struct{})
	for _, p := range k.Pairs.Iterate(ctx, collections.Range[string]{}).Keys() {
		pairsMap[p] = struct{}{}
	}

	return pairsMap
}

// resetExchangeRates removes all exchange rates from the state
func (k Keeper) resetExchangeRates(ctx sdk.Context) {
	for _, key := range k.ExchangeRates.Iterate(ctx, collections.Range[string]{}).Keys() {
		err := k.ExchangeRates.Delete(ctx, key)
		if err != nil {
			panic(err)
		}
	}
}

// getValidatorPerformanceMap returns a map [address]ValidatorPerformance excluding validators that are not bonded.
func (k Keeper) getValidatorPerformanceMap(ctx sdk.Context) map[string]types.ValidatorPerformance {
	validatorPerformanceMap := make(map[string]types.ValidatorPerformance)

	maxValidators := k.StakingKeeper.MaxValidators(ctx)
	powerReduction := k.StakingKeeper.PowerReduction(ctx)

	iterator := k.StakingKeeper.ValidatorsPowerStoreIterator(ctx)
	defer iterator.Close()

	i := 0
	for ; iterator.Valid() && i < int(maxValidators); iterator.Next() {
		validator := k.StakingKeeper.Validator(ctx, iterator.Value())

		// exclude not bonded
		if !validator.IsBonded() {
			continue
		}

		valAddr := validator.GetOperator()
		validatorPerformanceMap[valAddr.String()] = types.NewValidatorPerformance(validator.GetConsensusPower(powerReduction), 0, 0, valAddr)
		i++
	}

	return validatorPerformanceMap
}
