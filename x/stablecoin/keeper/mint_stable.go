// Package keeper for minting USDM  Minting USDM
// See Example B of https://docs.frax.finance/minting-and-redeeming
package keeper

import (
	"context"

	"github.com/MatrixDao/matrix/x/common"
	"github.com/MatrixDao/matrix/x/stablecoin/events"
	"github.com/MatrixDao/matrix/x/stablecoin/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MintStable mints stable coins given collateral and governance
// govDeposited: Units of GOV burned
// govDeposited = (1 - collRatio) * (collDeposited * 1) / (collRatio * priceGOV)
func (k Keeper) MintStable(
	goCtx context.Context, msg *types.MsgMintStable,
) (*types.MsgMintStableResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	msgCreator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// priceGov: Price of the governance token in USD
	priceGov, err := k.priceKeeper.GetCurrentPrice(ctx, common.GovPricePool)
	if err != nil {
		return nil, err
	}

	// priceColl: Price of the collateral token in USD
	priceColl, err := k.priceKeeper.GetCurrentPrice(ctx, common.CollPricePool)
	if err != nil {
		return nil, err
	}

	// The user deposits a mixture of collateral and GOV tokens based on the collateral ratio.
	// TODO: Initialize these two vars based on the collateral ratio of the protocol.
	collRatio, _ := sdk.NewDecFromStr("0.9")
	govRatio := sdk.OneDec().Sub(collRatio)

	neededCollUSD := sdk.NewDecFromInt(msg.Stable.Amount).Mul(collRatio)
	neededCollAmt := AsInt(neededCollUSD.Quo(priceColl.Price))
	neededColl := sdk.NewCoin(common.CollDenom, neededCollAmt)

	neededGovUSD := sdk.NewDecFromInt(msg.Stable.Amount).Mul(govRatio)
	neededGovAmt := AsInt(neededGovUSD.Quo(priceGov.Price))
	neededGov := sdk.NewCoin(common.GovDenom, neededGovAmt)

	coinsNeededToMint := sdk.NewCoins(neededColl, neededGov)

	err = k.CheckEnoughBalances(ctx, coinsNeededToMint, msgCreator)
	if err != nil {
		return nil, err
	}

	// Take assets out of the user account.
	err = k.bankKeeper.SendCoinsFromAccountToModule(
		ctx, msgCreator, types.ModuleName, coinsNeededToMint)
	if err != nil {
		panic(err)
	}

	for _, coin := range coinsNeededToMint {
		events.EmitTransfer(ctx, coin, msgCreator.String(), types.ModuleName)
	}

	err = k.mintStable(msg.Stable, ctx)
	if err != nil {
		panic(err)
	}

	err = k.burnGovTokens(ctx, neededGov)
	if err != nil {
		panic(err)
	}

	err = k.sendMintedTokensToUser(ctx, msgCreator, msg)
	if err != nil {
		panic(err)
	}

	return &types.MsgMintStableResponse{Stable: msg.Stable}, nil
}

func (k Keeper) sendMintedTokensToUser(ctx sdk.Context, msgCreator sdk.AccAddress, msg *types.MsgMintStable) error {
	err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, msgCreator, sdk.NewCoins(msg.Stable))
	if err != nil {
		return err
	}

	events.EmitTransfer(ctx, msg.Stable, types.ModuleName, msgCreator.String())

	return nil
}

func (k Keeper) burnGovTokens(ctx sdk.Context, neededGov sdk.Coin) error {
	err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(neededGov))
	if err != nil {
		return err
	}

	events.EmitBurnMtrx(ctx, neededGov)

	return nil
}

func (k Keeper) mintStable(stable sdk.Coin, ctx sdk.Context) error {
	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(stable))
	if err != nil {
		return err
	}

	events.EmitMintStable(ctx, stable)

	return nil
}
