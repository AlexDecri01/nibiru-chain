package testutil

import (
	"github.com/NibiruChain/nibiru/app"
	"github.com/NibiruChain/nibiru/simapp"
	"github.com/NibiruChain/nibiru/x/common"
	testutilcli "github.com/NibiruChain/nibiru/x/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestIntegrationTestSuite(t *testing.T) {
	coinsFromGenesis := []string{
		common.DenomGov,
		common.DenomStable,
		common.DenomColl,
		"coin-1",
		"coin-2",
		"coin-3",
		"coin-4",
		"coin-5",
	}

	app.SetPrefixes(app.AccountAddressPrefix)
	genesisState := simapp.NewTestGenesisStateFromDefault()

	genesisState = WhitelistGenesisAssets(
		genesisState,
		coinsFromGenesis,
	)

	cfg := testutilcli.BuildNetworkConfig(genesisState)
	cfg.StartingTokens = sdk.NewCoins(
		sdk.NewInt64Coin(common.DenomGov, 2e12), // for pool creation fee and more for tx fees
	)
	coinsToSendToUser := sdk.NewCoins(
		sdk.NewInt64Coin(common.DenomGov, 2e9), // for pool creation fee and more for tx fees
	)
	for _, coin := range coinsFromGenesis {
		cfg.StartingTokens = cfg.StartingTokens.Add(sdk.NewInt64Coin(coin, 40000))
		coinsToSendToUser = coinsToSendToUser.Add(sdk.NewInt64Coin(coin, 20000))
	}
	suite.Run(t, NewIntegrationTestSuite(cfg))
}
