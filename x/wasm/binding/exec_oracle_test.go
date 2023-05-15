package binding_test

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/NibiruChain/nibiru/app"
	"github.com/NibiruChain/nibiru/x/common/denoms"
	"github.com/NibiruChain/nibiru/x/common/testutil"
	"github.com/NibiruChain/nibiru/x/common/testutil/testapp"
	"github.com/NibiruChain/nibiru/x/wasm/binding"
	"github.com/NibiruChain/nibiru/x/wasm/binding/cw_struct"
	"github.com/NibiruChain/nibiru/x/wasm/binding/wasmbin"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"testing"
	"time"
)

func TestSuiteOracleExecutor_RunAll(t *testing.T) {
	suite.Run(t, new(TestSuiteOracleExecutor))
}

type TestSuiteOracleExecutor struct {
	suite.Suite

	nibiru           app.NibiruApp
	contractDeployer sdk.AccAddress
	exec             binding.ExecutorOracle
	contract         sdk.AccAddress
	ctx              sdk.Context
}

func (s *TestSuiteOracleExecutor) SetupSuite() {
	sender := testutil.AccAddress()
	s.contractDeployer = sender

	genesisState := SetupPerpGenesis()
	nibiru := testapp.NewNibiruTestApp(genesisState)
	ctx := nibiru.NewContext(false, tmproto.Header{
		Height:  1,
		ChainID: "nibiru-wasmnet-1",
		Time:    time.Now().UTC(),
	})

	coins := sdk.NewCoins(
		sdk.NewCoin(denoms.NIBI, sdk.TokensFromConsensusPower(1, sdk.DefaultPowerReduction)),
		sdk.NewCoin(denoms.NUSD, sdk.TokensFromConsensusPower(1, sdk.DefaultPowerReduction)),
	)
	s.NoError(testapp.FundAccount(nibiru.BankKeeper, ctx, sender, coins))

	nibiru, ctx = SetupAllContracts(s.T(), sender, nibiru, ctx)
	s.nibiru = *nibiru
	s.ctx = ctx

	wasmkeeper.NewMsgServerImpl(wasmkeeper.NewDefaultPermissionKeeper(nibiru.WasmKeeper))
	s.contract = ContractMap[wasmbin.WasmKeyPerpBinding]
	s.exec = binding.ExecutorOracle{
		Oracle: nibiru.OracleKeeper,
	}
}

func (s *TestSuiteOracleExecutor) TestExecuteOracleParams() {
	period := uint64(1234)
	cwMsg := cw_struct.BindingMsg{
		OracleParams: &cw_struct.OracleParams{
			VotePeriod:         &period,
			VoteThreshold:      nil,
			RewardBand:         nil,
			Whitelist:          nil,
			SlashFraction:      nil,
			SlashWindow:        nil,
			MinValidPerWindow:  nil,
			TwapLookbackWindow: nil,
			MinVoters:          nil,
			ValidatorFeeRatio:  nil,
		},
	}

	bz, err := DoCustomBindingExecute(s.ctx, &s.nibiru, s.contract, s.contractDeployer, cwMsg, sdk.Coins{})
	s.NoError(err)
	s.NotNil(bz)
}
