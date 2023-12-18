package action

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/NibiruChain/nibiru/app"
	"github.com/NibiruChain/nibiru/x/common/asset"
	"github.com/NibiruChain/nibiru/x/common/testutil/action"
	"github.com/NibiruChain/nibiru/x/perp/v2/types"
)

type PairTraderTuple struct {
	Pair       asset.Pair
	Trader     sdk.AccAddress
	Successful bool
}

type multiLiquidate struct {
	pairTraderTuples []PairTraderTuple
	liquidator       sdk.AccAddress
	shouldAllFail    bool
}

func (m multiLiquidate) Do(app *app.NibiruApp, ctx sdk.Context) (sdk.Context, error) {
	liquidationRequests := make([]*types.MsgMultiLiquidate_Liquidation, len(m.pairTraderTuples))

	for i, pairTraderTuple := range m.pairTraderTuples {
		liquidationRequests[i] = &types.MsgMultiLiquidate_Liquidation{
			Pair:   pairTraderTuple.Pair,
			Trader: pairTraderTuple.Trader.String(),
		}
	}

	responses, err := app.PerpKeeperV2.MultiLiquidate(ctx, m.liquidator, liquidationRequests)

	if m.shouldAllFail {
		// we check if all liquidations failed
		if err == nil {
			return ctx, fmt.Errorf("multi liquidations should have all failed, but instead some succeeded")
		}

		for i, response := range responses {
			if response.Success {
				return ctx, fmt.Errorf("multi liquidations should have all failed, but instead some succeeded, index %d", i)
			}
		}

		return ctx, nil
	}

	// otherwise, some succeeded and some may have failed
	if err != nil {
		return ctx, err
	}

	for i, response := range responses {
		if response.Success != m.pairTraderTuples[i].Successful {
			return ctx, fmt.Errorf("MultiLiquidate wrong assertion, expected %v, got %v, index %d", m.pairTraderTuples[i].Successful, response.Success, i)
		}
	}

	return ctx, nil
}

func MultiLiquidate(liquidator sdk.AccAddress, shouldAllFail bool, pairTraderTuples ...PairTraderTuple) action.Action {
	return multiLiquidate{
		pairTraderTuples: pairTraderTuples,
		liquidator:       liquidator,
		shouldAllFail:    shouldAllFail,
	}
}
