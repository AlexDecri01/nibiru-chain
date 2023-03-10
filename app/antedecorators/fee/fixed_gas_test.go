package fee_test

import (
	"github.com/NibiruChain/nibiru/app"
	"github.com/NibiruChain/nibiru/app/antedecorators/fee"
	oracletypes "github.com/NibiruChain/nibiru/x/oracle/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"
)

func (suite *AnteTestSuite) TestOraclePostPriceTransactionsHaveFixedPrice() {
	priv1, _, addr := testdata.KeyTestPubAddr()

	tests := []struct {
		name        string
		messages    []sdk.Msg
		expectedGas sdk.Gas
		expectedErr error
	}{
		{
			name: "Oracle Prevote Transaction",
			messages: []sdk.Msg{
				&oracletypes.MsgAggregateExchangeRatePrevote{
					Hash:      "dummyData",
					Feeder:    addr.String(),
					Validator: addr.String(),
				},
			},
			expectedGas: fee.OracleMessageGas,
			expectedErr: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		suite.T().Run(tc.name, func(t *testing.T) {
			suite.SetupTest() // setup
			suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

			// msg and signatures
			feeAmount := sdk.NewCoins(sdk.NewInt64Coin(app.BondDenom, 150))
			gasLimit := testdata.NewTestGasLimit()
			suite.txBuilder.SetFeeAmount(feeAmount)
			suite.txBuilder.SetGasLimit(gasLimit)
			suite.Require().NoError(suite.txBuilder.SetMsgs(tc.messages...))

			privs, accNums, accSeqs := []cryptotypes.PrivKey{priv1}, []uint64{12}, []uint64{0}
			tx, err := suite.CreateTestTx(privs, accNums, accSeqs, suite.ctx.ChainID())
			suite.Require().NoError(err)

			err = simapp.FundAccount(suite.app.BankKeeper, suite.ctx, addr, sdk.NewCoins(sdk.NewInt64Coin(app.BondDenom, 1000)))
			if err != nil {
				return
			}

			suite.ctx, err = suite.anteHandler(suite.ctx, tx, false)
			suite.Require().Equal(tc.expectedErr, err)
			suite.Require().Equal(tc.expectedGas, suite.ctx.GasMeter().GasConsumed())
		})
	}
}
