package keeper

import (
	"context"

	"github.com/NibiruChain/nibiru/collections/keys"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NibiruChain/nibiru/x/common"
	"github.com/NibiruChain/nibiru/x/perp/types"
)

type queryServer struct {
	k Keeper
}

func NewQuerier(k Keeper) types.QueryServer {
	return queryServer{k: k}
}

var _ types.QueryServer = queryServer{}

func (q queryServer) QueryTraderPosition(
	goCtx context.Context, req *types.QueryTraderPositionRequest,
) (*types.QueryTraderPositionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	_, err := sdk.AccAddressFromBech32(req.Trader) // just for validation purposes
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	pair, err := common.NewAssetPair(req.TokenPair)
	if err != nil {
		return nil, err
	}

	return q.traderPosition(ctx, pair, req.Trader)
}

func (q queryServer) QueryTraderPositions(goCtx context.Context, req *types.QueryTraderPositionsRequest) (*types.QueryTraderPositionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	rng := keys.NewRange[keys.Pair[keys.StringKey, keys.Pair[common.AssetPair, keys.StringKey]]]()
	rng = rng.Prefix(keys.PairPrefix[keys.StringKey, keys.Pair[common.AssetPair, keys.StringKey]](keys.String(req.Trader)))

	pks := q.k.Positions.Indexes.Address.Iterate(ctx, rng).Keys()
	positions := make([]*types.QueryTraderPositionResponse, len(pks))
	for i, pk := range pks {
		posResp, err := q.traderPosition(ctx, pk.K1(), string(pk.K2()))
		if err != nil {
			return nil, err
		}
		positions[i] = posResp
	}

	return &types.QueryTraderPositionsResponse{Positions: positions}, nil
}

func (q queryServer) traderPosition(ctx sdk.Context, pair common.AssetPair, trader string) (*types.QueryTraderPositionResponse, error) {
	position, err := q.k.Positions.Get(ctx, keys.Join(pair, keys.String(trader)))
	if err != nil {
		return nil, err
	}

	positionNotional, unrealizedPnl, err := q.k.getPositionNotionalAndUnrealizedPnL(ctx, position, types.PnLCalcOption_SPOT_PRICE)
	if err != nil {
		return nil, err
	}

	marginRatioMark, err := q.k.GetMarginRatio(ctx, position, types.MarginCalculationPriceOption_MAX_PNL)
	if err != nil {
		return nil, err
	}
	marginRatioIndex, err := q.k.GetMarginRatio(ctx, position, types.MarginCalculationPriceOption_INDEX)
	if err != nil {
		// The index portion of the query fails silently as not to distrupt all
		// position queries when oracles aren't posting prices.
		q.k.Logger(ctx).Error(err.Error())
		marginRatioIndex = sdk.Dec{}
	}

	return &types.QueryTraderPositionResponse{
		Position:         &position,
		PositionNotional: positionNotional,
		UnrealizedPnl:    unrealizedPnl,
		MarginRatioMark:  marginRatioMark,
		MarginRatioIndex: marginRatioIndex,
		BlockNumber:      ctx.BlockHeight(),
	}, nil
}

func (q queryServer) Params(
	goCtx context.Context, req *types.QueryParamsRequest,
) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryParamsResponse{Params: q.k.GetParams(ctx)}, nil
}

func (q queryServer) FundingRates(
	goCtx context.Context, req *types.QueryFundingRatesRequest,
) (*types.QueryFundingRatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	assetPair, err := common.NewAssetPair(req.Pair)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid pair: %s", req.Pair)
	}

	pairMetadata, err := q.k.PairsMetadata.Get(ctx, assetPair)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "could not find pair: %s", req.Pair)
	}

	var fundingRates []sdk.Dec

	// truncate to most recent 48 funding payments
	// given 30 minute funding rate calculations, this should give the last 24 hours of funding payments
	numFundingRates := len(pairMetadata.CumulativeFundingRates)
	if numFundingRates >= 48 {
		fundingRates = pairMetadata.CumulativeFundingRates[numFundingRates-48 : numFundingRates]
	} else {
		fundingRates = pairMetadata.CumulativeFundingRates
	}

	return &types.QueryFundingRatesResponse{
		CumulativeFundingRates: fundingRates,
	}, nil
}
