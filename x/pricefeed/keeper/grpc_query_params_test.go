package keeper_test

import (
	"testing"

	"github.com/NibiruChain/nibiru/x/pricefeed/types"
	"github.com/NibiruChain/nibiru/x/testutil/testkeeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.PricefeedKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.Params{
		Pairs: []types.Pair{
			{PairID: "btc:usd", BaseAsset: "btc", QuoteAsset: "usd", Oracles: nil, Active: true},
			{PairID: "xrp:usd", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: nil, Active: true},
			{PairID: "xrp:usd:30", BaseAsset: "xrp", QuoteAsset: "usd", Oracles: nil, Active: true},
		},
	}
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
