package keeper_test

import (
	"testing"
	"time"

	"github.com/NibiruChain/nibiru/x/common"
	pricefeedTypes "github.com/NibiruChain/nibiru/x/pricefeed/types"
	"github.com/NibiruChain/nibiru/x/stablecoin/types"
	"github.com/NibiruChain/nibiru/x/testutil"
	"github.com/NibiruChain/nibiru/x/testutil/sample"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestSetCollRatio_Input(t *testing.T) {
	type TestCase struct {
		name         string
		inCollRatio  sdk.Dec
		expectedPass bool
	}

	executeTest := func(t *testing.T, testCase TestCase) {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			nibiruApp, ctx := testutil.NewNibiruApp(true)
			stablecoinKeeper := &nibiruApp.StablecoinKeeper

			err := stablecoinKeeper.SetCollRatio(ctx, tc.inCollRatio)
			if tc.expectedPass {
				require.NoError(
					t, err, "Error setting the CollRatio: %d", tc.inCollRatio)
				return
			}
			require.Error(t, err)
		})
	}

	testCases := []TestCase{
		{
			name:         "Upper bound of CollRatio",
			inCollRatio:  sdk.OneDec(),
			expectedPass: true,
		}, {
			name:         "Lower bound of CollRatio",
			inCollRatio:  sdk.ZeroDec(),
			expectedPass: true,
		}, {
			name:         "CollRatio above 100",
			inCollRatio:  sdk.MustNewDecFromStr("1.5"),
			expectedPass: false,
		}, {
			name:         "Negative CollRatio not allowed",
			inCollRatio:  sdk.OneDec().Neg(),
			expectedPass: false,
		},
	}
	for _, testCase := range testCases {
		executeTest(t, testCase)
	}
}

func TestGetCollRatio_Input(t *testing.T) {
	testName := "GetCollRatio after setting default params returns expected value"
	t.Run(testName, func(t *testing.T) {
		nibiruApp, ctx := testutil.NewNibiruApp(true)
		stablecoinKeeper := &nibiruApp.StablecoinKeeper

		stablecoinKeeper.SetParams(ctx, types.DefaultParams())
		expectedCollRatioInt := sdk.NewInt(types.DefaultParams().CollRatio)

		outCollRatio := stablecoinKeeper.GetCollRatio(ctx)
		outCollRatioInt := outCollRatio.Mul(sdk.MustNewDecFromStr("1000000")).RoundInt()
		require.EqualValues(t, expectedCollRatioInt, outCollRatioInt)
	})

	testName = "Setting to non-default value returns expected value"
	t.Run(testName, func(t *testing.T) {
		nibiruApp, ctx := testutil.NewNibiruApp(true)
		stablecoinKeeper := &nibiruApp.StablecoinKeeper

		expectedCollRatio := sdk.MustNewDecFromStr("0.5")
		expectedCollRatioInt := expectedCollRatio.Mul(sdk.MustNewDecFromStr("1000000")).RoundInt()
		require.NoError(t, stablecoinKeeper.SetCollRatio(ctx, expectedCollRatio))

		outCollRatio := stablecoinKeeper.GetCollRatio(ctx)
		outCollRatioInt := outCollRatio.Mul(sdk.MustNewDecFromStr("1000000")).RoundInt()
		require.EqualValues(t, expectedCollRatioInt, outCollRatioInt)
	})
}

func TestGetUSDValForTargetCollRatio(t *testing.T) {
	testCases := []struct {
		name             string
		protocolColl     sdk.Int
		priceCollStable  sdk.Dec
		postedAssetPairs []common.AssetPair
		stableSupply     sdk.Int
		targetCollRatio  sdk.Dec
		neededUSD        sdk.Dec

		expectedPass bool
	}{
		{
			name:            "Too little collateral gives correct positive value",
			protocolColl:    sdk.NewInt(500),
			priceCollStable: sdk.OneDec(), // startCollUSD = 500 * 1 -> 500
			postedAssetPairs: []common.AssetPair{
				common.CollStablePool,
			},
			stableSupply:    sdk.NewInt(1000),
			targetCollRatio: sdk.MustNewDecFromStr("0.6"), // 0.6 * 1000 = 600
			neededUSD:       sdk.MustNewDecFromStr("100"), // = 600 - 500
			expectedPass:    true,
		}, {
			name:            "Too much collateral gives correct negative value",
			protocolColl:    sdk.NewInt(600),
			priceCollStable: sdk.OneDec(), // startCollUSD = 600 * 1 = 600
			postedAssetPairs: []common.AssetPair{
				common.CollStablePool,
			},
			stableSupply:    sdk.NewInt(1000),
			targetCollRatio: sdk.MustNewDecFromStr("0.5"),  // 0.5 * 1000 = 500
			neededUSD:       sdk.MustNewDecFromStr("-100"), // = 500 - 600
			expectedPass:    true,
		}, {
			name:             "No price availabale for the collateral",
			protocolColl:     sdk.NewInt(500),
			priceCollStable:  sdk.OneDec(), // startCollUSD = 500 * 1 -> 500
			postedAssetPairs: []common.AssetPair{},
			stableSupply:     sdk.NewInt(1_000),
			targetCollRatio:  sdk.MustNewDecFromStr("0.6"), // 0.6 * 1000 = 600
			neededUSD:        sdk.MustNewDecFromStr("100"), // = 600 - 500
			expectedPass:     false,
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			nibiruApp, ctx := testutil.NewNibiruApp(true)
			stablecoinKeeper := &nibiruApp.StablecoinKeeper
			require.NoError(t, stablecoinKeeper.SetCollRatio(ctx, tc.targetCollRatio))
			require.NoError(t, nibiruApp.BankKeeper.MintCoins(
				ctx, types.ModuleName, sdk.NewCoins(
					sdk.NewCoin(common.CollDenom, tc.protocolColl),
					sdk.NewCoin(common.StableDenom, tc.stableSupply),
				),
			))

			// Set up markets for the pricefeed keeper.
			oracle := sample.AccAddress()
			priceExpiry := ctx.BlockTime().Add(time.Hour)
			pricefeedParams := pricefeedTypes.Params{
				Pairs: []pricefeedTypes.Pair{
					{Token0: common.CollDenom,
						Token1:  common.StableDenom,
						Oracles: []sdk.AccAddress{oracle}, Active: true},
					{Token0: common.GovDenom,
						Token1:  common.StableDenom,
						Oracles: []sdk.AccAddress{oracle}, Active: true},
				}}
			nibiruApp.PriceKeeper.SetParams(ctx, pricefeedParams)

			// Post prices to each specified market with the oracle.
			prices := map[common.AssetPair]sdk.Dec{
				common.CollStablePool: tc.priceCollStable,
			}
			for _, pair := range tc.postedAssetPairs {
				_, err := nibiruApp.PriceKeeper.SetPrice(
					ctx, oracle, pair.Token0, pair.Token1, prices[pair], priceExpiry)
				require.NoError(t, err)

				// Update the 'CurrentPrice' posted by the oracles.
				err = nibiruApp.PriceKeeper.SetCurrentPrices(ctx, pair.Token0, pair.Token1)
				require.NoError(t, err, "Error posting price for pair: %d", pair.String())
			}

			neededUSD, err := stablecoinKeeper.GetUSDValForTargetCollRatio(ctx)
			if tc.expectedPass {
				require.NoError(t, err)
				require.EqualValues(t, tc.neededUSD, neededUSD)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestRecollateralizeCollAmtForTargetCollRatio(t *testing.T) {
	type TestCaseRecollateralizeCollAmtForTargetCollRatio struct {
		name            string
		protocolColl    sdk.Int
		priceCollStable sdk.Dec
		stableSupply    sdk.Int
		targetCollRatio sdk.Dec
		neededCollAmt   sdk.Int
		expectedPass    bool
	}

	expectedPasses := []TestCaseRecollateralizeCollAmtForTargetCollRatio{
		{
			name:            "under-collateralized; untruncated integer amount",
			protocolColl:    sdk.NewInt(500),
			priceCollStable: sdk.OneDec(), // startCollUSD = 500 * 1 -> 500
			stableSupply:    sdk.NewInt(1000),
			targetCollRatio: sdk.MustNewDecFromStr("0.6"), // 0.6 * 1000 = 600
			neededCollAmt:   sdk.NewInt(100),              // = 600 - 500
			expectedPass:    true,
		},
		{
			name:            "under-collateralized; truncated integer amount",
			protocolColl:    sdk.NewInt(500),
			priceCollStable: sdk.OneDec(), // startCollUSD = 500 * 1 -> 500
			stableSupply:    sdk.NewInt(1000),
			targetCollRatio: sdk.MustNewDecFromStr("0.6025"), // 0.6025 * 1000 = 602.5
			neededCollAmt:   sdk.NewInt(103),                 //  602.5 - 500 -> 103 required
			expectedPass:    true,
		},
		{
			name:            "under-collateralized; truncated integer amount; non-unit price",
			protocolColl:    sdk.NewInt(500),
			priceCollStable: sdk.MustNewDecFromStr("0.999"), // startCollUSD = 500 * 0.999 -> 499.5
			stableSupply:    sdk.NewInt(1000),
			targetCollRatio: sdk.MustNewDecFromStr("0.603"), // 0.603 * 1000 = 603
			//  603 - 499.5 = 103.5 -> 104 required
			neededCollAmt: sdk.NewInt(104),
			expectedPass:  true,
		},
	}

	for _, testCase := range expectedPasses {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			nibiruApp, ctx := testutil.NewNibiruApp(true)
			stablecoinKeeper := &nibiruApp.StablecoinKeeper
			require.NoError(t, stablecoinKeeper.SetCollRatio(ctx, tc.targetCollRatio))
			require.NoError(t, nibiruApp.BankKeeper.MintCoins(
				ctx, types.ModuleName, sdk.NewCoins(
					sdk.NewCoin(common.CollDenom, tc.protocolColl),
					sdk.NewCoin(common.StableDenom, tc.stableSupply),
				),
			))

			// Set up markets for the pricefeed keeper.
			pair := common.CollStablePool
			oracle := sample.AccAddress()
			priceExpiry := ctx.BlockTime().Add(time.Hour)
			pricefeedParams := pricefeedTypes.Params{
				Pairs: []pricefeedTypes.Pair{
					{Token1: common.CollDenom,
						Token0:  common.StableDenom,
						Oracles: []sdk.AccAddress{oracle}, Active: true},
				}}
			nibiruApp.PriceKeeper.SetParams(ctx, pricefeedParams)

			// Post prices to each market with the oracle.
			_, err := nibiruApp.PriceKeeper.SetPrice(
				ctx, oracle, pair.Token0, pair.Token1, tc.priceCollStable, priceExpiry)
			require.NoError(t, err)

			// Update the 'CurrentPrice' posted by the oracles.
			for _, pfPair := range pricefeedParams.Pairs {
				err = nibiruApp.PriceKeeper.SetCurrentPrices(ctx, pfPair.Token0, pfPair.Token1)
				require.NoError(t, err, "Error posting price for market: %d", pfPair.AsString())
			}

			neededCollAmount, err := stablecoinKeeper.RecollateralizeCollAmtForTargetCollRatio(ctx)
			if tc.expectedPass {
				require.NoError(t, err)
				require.EqualValues(t, tc.neededCollAmt, neededCollAmount)
			} else {
				require.Error(t, err)
			}
		})
	}

	expectedFails := []TestCaseRecollateralizeCollAmtForTargetCollRatio{
		{
			name:            "error from price not being posted",
			protocolColl:    sdk.NewInt(500),
			priceCollStable: sdk.OneDec(), // startCollUSD = 500 * 1 -> 500
			stableSupply:    sdk.NewInt(1000),
			targetCollRatio: sdk.MustNewDecFromStr("0.6"), // 0.6 * 1000 = 600
			neededCollAmt:   sdk.NewInt(100),              // = 600 - 500
			expectedPass:    false,
		},
	}

	for _, testCase := range expectedFails {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			nibiruApp, ctx := testutil.NewNibiruApp(true)
			stablecoinKeeper := &nibiruApp.StablecoinKeeper
			require.NoError(t, stablecoinKeeper.SetCollRatio(ctx, tc.targetCollRatio))
			require.NoError(t, nibiruApp.BankKeeper.MintCoins(
				ctx, types.ModuleName, sdk.NewCoins(
					sdk.NewCoin(common.CollDenom, tc.protocolColl),
					sdk.NewCoin(common.StableDenom, tc.stableSupply),
				),
			))

			// Set up markets for the pricefeed keeper.
			oracle := sample.AccAddress()
			pricefeedParams := pricefeedTypes.Params{
				Pairs: []pricefeedTypes.Pair{
					{Token1: common.CollDenom,
						Token0:  common.StableDenom,
						Oracles: []sdk.AccAddress{oracle}, Active: true},
				}}
			nibiruApp.PriceKeeper.SetParams(ctx, pricefeedParams)

			neededCollAmount, err := stablecoinKeeper.RecollateralizeCollAmtForTargetCollRatio(ctx)
			if tc.expectedPass {
				require.NoError(t, err)
				require.EqualValues(t, tc.neededCollAmt, neededCollAmount)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestGovAmtFromFullRecollateralize(t *testing.T) {
	testCases := []struct {
		name             string
		protocolColl     sdk.Int
		priceCollStable  sdk.Dec
		priceGovStable   sdk.Dec
		stableSupply     sdk.Int
		targetCollRatio  sdk.Dec
		postedAssetPairs []common.AssetPair

		govOut       sdk.Int
		expectedPass bool
	}{
		{
			name:             "no prices posted",
			protocolColl:     sdk.NewInt(500),
			stableSupply:     sdk.NewInt(1000),
			targetCollRatio:  sdk.MustNewDecFromStr("0.6"),
			postedAssetPairs: []common.AssetPair{},
			govOut:           sdk.Int{},
			expectedPass:     false,
		},
		{
			name:            "only post collateral price",
			protocolColl:    sdk.NewInt(500),
			stableSupply:    sdk.NewInt(1000),
			targetCollRatio: sdk.MustNewDecFromStr("0.6"), // 0.6 * 1000 = 600
			priceCollStable: sdk.OneDec(),
			postedAssetPairs: []common.AssetPair{
				common.CollStablePool},
			govOut:       sdk.Int{},
			expectedPass: false,
		},
		{
			name:            "only post gov price",
			protocolColl:    sdk.NewInt(500),
			stableSupply:    sdk.NewInt(1000),
			targetCollRatio: sdk.MustNewDecFromStr("0.6"), // 0.6 * 1000 = 600
			priceGovStable:  sdk.OneDec(),
			postedAssetPairs: []common.AssetPair{
				common.GovStablePool},
			govOut:       sdk.Int{},
			expectedPass: false,
		},
		{
			name:            "correct computation - positive",
			protocolColl:    sdk.NewInt(5_000),
			stableSupply:    sdk.NewInt(10_000),
			targetCollRatio: sdk.MustNewDecFromStr("0.7"), // 0.7 * 10_000 = 7_000
			postedAssetPairs: []common.AssetPair{
				common.GovStablePool,
				common.CollStablePool,
			},
			priceCollStable: sdk.OneDec(),
			priceGovStable:  sdk.NewDec(2),
			// govOut = neededUSD * (1 + bonusRate) / priceGov
			//        = 2000 * (1.002) / 2 = 1002
			govOut:       sdk.NewInt(1002),
			expectedPass: true,
		},
		{
			name:            "correct computation - positive, new price",
			protocolColl:    sdk.NewInt(50_000),
			stableSupply:    sdk.NewInt(100_000),
			targetCollRatio: sdk.MustNewDecFromStr("0.7"), // 0.7 * 100_000 = 70_000
			postedAssetPairs: []common.AssetPair{
				common.GovStablePool,
				common.CollStablePool,
			},
			priceCollStable: sdk.OneDec(),
			priceGovStable:  sdk.NewDec(10),
			// govOut = neededUSD * (1 + bonusRate) / priceGov
			//        = 20000 * (1.002) / 10 = 2004
			govOut:       sdk.NewInt(2004),
			expectedPass: true,
		},
		{
			name:            "correct computation - negative",
			protocolColl:    sdk.NewInt(70_000),
			stableSupply:    sdk.NewInt(100_000),
			targetCollRatio: sdk.MustNewDecFromStr("0.5"), // 0.5 * 100_000 = 50_000
			postedAssetPairs: []common.AssetPair{
				common.GovStablePool,
				common.CollStablePool,
			},
			priceCollStable: sdk.OneDec(),
			priceGovStable:  sdk.NewDec(10),
			// govOut = neededUSD * (1 + bonusRate) / priceGov
			//        = -20000 * (1.002) / 10 = 2004
			govOut:       sdk.NewInt(-2004),
			expectedPass: true,
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			nibiruApp, ctx := testutil.NewNibiruApp(true)
			stablecoinKeeper := &nibiruApp.StablecoinKeeper
			require.NoError(t, stablecoinKeeper.SetCollRatio(ctx, tc.targetCollRatio))
			require.NoError(t, nibiruApp.BankKeeper.MintCoins(
				ctx, types.ModuleName, sdk.NewCoins(
					sdk.NewCoin(common.CollDenom, tc.protocolColl),
					sdk.NewCoin(common.StableDenom, tc.stableSupply),
				),
			))

			// Set up markets for the pricefeed keeper.
			oracle := sample.AccAddress()
			priceExpiry := ctx.BlockTime().Add(time.Hour)
			pricefeedParams := pricefeedTypes.Params{
				Pairs: []pricefeedTypes.Pair{
					{Token1: common.CollDenom,
						Token0:  common.StableDenom,
						Oracles: []sdk.AccAddress{oracle}, Active: true},
					{Token1: common.GovDenom,
						Token0:  common.StableDenom,
						Oracles: []sdk.AccAddress{oracle}, Active: true},
				}}
			nibiruApp.PriceKeeper.SetParams(ctx, pricefeedParams)

			prices := map[common.AssetPair]sdk.Dec{
				common.GovStablePool:  tc.priceGovStable,
				common.CollStablePool: tc.priceCollStable,
			}
			for _, pair := range tc.postedAssetPairs {
				_, err := nibiruApp.PriceKeeper.SetPrice(
					ctx, oracle, pair.Token0, pair.Token1, prices[pair], priceExpiry)
				require.NoError(t, err)

				// Update the 'CurrentPrice' posted by the oracles.
				err = nibiruApp.PriceKeeper.SetCurrentPrices(ctx, pair.Token0, pair.Token1)
				require.NoError(t, err, "Error posting price for pair: %d", pair.String())
			}

			// Post prices to each specified market with the oracle.
			prices = map[common.AssetPair]sdk.Dec{
				common.CollStablePool: tc.priceCollStable,
				common.GovStablePool:  tc.priceGovStable,
			}
			for _, assetPair := range tc.postedAssetPairs {
				_, err := nibiruApp.PriceKeeper.SetPrice(
					ctx, oracle, assetPair.Token0, assetPair.Token1,
					prices[assetPair], priceExpiry)
				require.NoError(t, err)

				// Update the 'CurrentPrice' posted by the oracles.
				err = nibiruApp.PriceKeeper.SetCurrentPrices(
					ctx, assetPair.Token0, assetPair.Token1)
				require.NoError(
					t, err, "Error posting price for pair: %d", assetPair.String())
			}

			govOut, err := stablecoinKeeper.GovAmtFromFullRecollateralize(ctx)
			if tc.expectedPass {
				require.NoError(t, err)
				require.EqualValues(t, tc.govOut, govOut)
			} else {
				require.Error(t, err)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Buyback and Recollateralize Tests
// ---------------------------------------------------------------------------

type NeededCollScenario struct {
	protocolColl    sdk.Int
	priceCollStable sdk.Dec
	stableSupply    sdk.Int
	collRatio       sdk.Dec
}

func (scenario NeededCollScenario) CalcNeededUSD() (neededUSD sdk.Dec) {
	stableUSD := scenario.collRatio.MulInt(scenario.stableSupply)
	collUSD := scenario.priceCollStable.MulInt(scenario.protocolColl)
	return stableUSD.Sub(collUSD)
}

func TestRecollateralize(t *testing.T) {
	testCases := []struct {
		name         string
		expectedPass bool

		postedAssetPairs  []common.AssetPair
		scenario          NeededCollScenario
		priceGovStable    sdk.Dec
		expectedNeededUSD sdk.Dec
		accFunds          sdk.Coins

		msg      types.MsgRecollateralize
		response *types.MsgRecollateralizeResponse
	}{
		{
			name: "both prices are $1",
			postedAssetPairs: []common.AssetPair{
				common.GovStablePool,
				common.CollStablePool,
			},
			scenario: NeededCollScenario{
				protocolColl:    sdk.NewInt(500_000),
				priceCollStable: sdk.OneDec(),
				stableSupply:    sdk.NewInt(1_000_000),
				collRatio:       sdk.MustNewDecFromStr("0.6"),
				// neededUSD =  (0.6 * 1000e3) - (500e3 *1) = 100_000
			},
			priceGovStable: sdk.OneDec(),
			accFunds: sdk.NewCoins(
				sdk.NewInt64Coin(common.CollDenom, 1_000_000_000),
			),

			expectedNeededUSD: sdk.NewDec(100_000),
			msg: types.MsgRecollateralize{
				Creator: sample.AccAddress().String(),
				Coll:    sdk.NewCoin(common.CollDenom, sdk.NewInt(100_000)),
			},
			response: &types.MsgRecollateralizeResponse{
				/*
					Gov.Amount = inCollUSD * (1 + bonusRate) / priceGovStable
					  = 100_000 * (1.002) / priceGovStable
					  = 100_200 / priceGovStable
				*/
				Gov: sdk.NewCoin(common.GovDenom, sdk.NewInt(100_200)),
			},
			expectedPass: true,
		},
		{
			name: "arbitrary valid prices",
			postedAssetPairs: []common.AssetPair{
				common.GovStablePool,
				common.CollStablePool,
			},
			scenario: NeededCollScenario{
				protocolColl:    sdk.NewInt(500_000),
				priceCollStable: sdk.MustNewDecFromStr("1.099999"),
				stableSupply:    sdk.NewInt(1_000_000),
				collRatio:       sdk.MustNewDecFromStr("0.7"),
				// neededUSD =  (0.7 * 1000e3) - (500e3 *1.09999) = 150_000.5
			},
			priceGovStable: sdk.NewDec(5),
			accFunds: sdk.NewCoins(
				sdk.NewInt64Coin(common.CollDenom, 1_000_000_000),
			),

			// Since 'neededUSD' is
			expectedNeededUSD: sdk.MustNewDecFromStr("150000.5"),
			msg: types.MsgRecollateralize{
				Creator: sample.AccAddress().String(),
				Coll:    sdk.NewCoin(common.CollDenom, sdk.NewInt(50_000)),
			},
			response: &types.MsgRecollateralizeResponse{
				/*
					Gov.Amount = inCollUSD * (1 + bonusRate) / priceGovStable
					  = msg.Coll.Amount * priceCollStable (1.002) / priceGovStable
					  = 50_000 * 1.099999 * (1.002) / priceGovStable
					  = 55109.9499 / priceGovStable
					  = 11021.98998 -> 11_021
				*/
				Gov: sdk.NewCoin(common.GovDenom, sdk.NewInt(11_021)),
			},
			expectedPass: true,
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			require.EqualValues(t, tc.expectedNeededUSD, tc.scenario.CalcNeededUSD())

			nibiruApp, ctx := testutil.NewNibiruApp(true)
			stablecoinKeeper := &nibiruApp.StablecoinKeeper
			require.NoError(t, stablecoinKeeper.SetCollRatio(ctx, tc.scenario.collRatio))
			require.NoError(t, nibiruApp.BankKeeper.MintCoins(
				ctx, types.ModuleName, sdk.NewCoins(
					sdk.NewCoin(common.CollDenom, tc.scenario.protocolColl),
					sdk.NewCoin(common.StableDenom, tc.scenario.stableSupply),
				),
			))
			// Fund account
			caller, err := sdk.AccAddressFromBech32(tc.msg.Creator)
			if tc.expectedPass {
				require.NoError(t, err)
			}
			err = simapp.FundAccount(nibiruApp.BankKeeper, ctx, caller, tc.accFunds)
			if tc.expectedPass {
				require.NoError(t, err)
			}

			// Set up markets for the pricefeed keeper.
			oracle := sample.AccAddress()
			priceExpiry := ctx.BlockTime().Add(time.Hour)
			pricefeedParams := pricefeedTypes.Params{
				Pairs: []pricefeedTypes.Pair{
					{Token0: common.CollDenom, Token1: common.StableDenom,
						Oracles: []sdk.AccAddress{oracle}, Active: true},
					{Token0: common.GovDenom, Token1: common.StableDenom,
						Oracles: []sdk.AccAddress{oracle}, Active: true},
				}}
			nibiruApp.PriceKeeper.SetParams(ctx, pricefeedParams)

			prices := map[common.AssetPair]sdk.Dec{
				common.GovStablePool:  tc.priceGovStable,
				common.CollStablePool: tc.scenario.priceCollStable,
			}
			for _, pair := range tc.postedAssetPairs {
				_, err := nibiruApp.PriceKeeper.SetPrice(
					ctx, oracle, pair.Token0, pair.Token1, prices[pair], priceExpiry)
				require.NoError(t, err)

				// Update the 'CurrentPrice' posted by the oracles.
				err = nibiruApp.PriceKeeper.SetCurrentPrices(ctx, pair.Token0, pair.Token1)
				require.NoError(t, err, "Error posting price for pair: %d", pair.String())
			}

			// Post prices to each specified market with the oracle.
			prices = map[common.AssetPair]sdk.Dec{
				common.CollStablePool: tc.scenario.priceCollStable,
				common.GovStablePool:  tc.priceGovStable,
			}
			for _, assetPair := range tc.postedAssetPairs {
				_, err := nibiruApp.PriceKeeper.SetPrice(
					ctx, oracle, assetPair.Token0, assetPair.Token1,
					prices[assetPair], priceExpiry)
				require.NoError(t, err)

				// Update the 'CurrentPrice' posted by the oracles.
				err = nibiruApp.PriceKeeper.SetCurrentPrices(
					ctx, assetPair.Token0, assetPair.Token1)
				require.NoError(
					t, err, "Error posting price for pair: %d", assetPair.String())
			}

			goCtx := sdk.WrapSDKContext(ctx)
			response, err := stablecoinKeeper.Recollateralize(goCtx, &tc.msg)
			if tc.expectedPass {
				require.NoError(t, err)
				require.EqualValues(t, tc.response, response)
			} else {
				require.Error(t, err)
			}
		},
		)
	}
}

func TestBuyback(t *testing.T) {
	testCases := []struct {
		name         string
		expectedPass bool

		postedAssetPairs  []common.AssetPair
		scenario          NeededCollScenario
		priceGovStable    sdk.Dec
		expectedNeededUSD sdk.Dec
		accFunds          sdk.Coins
		moduleAccFunds    sdk.Coins

		msg      types.MsgBuyback
		response *types.MsgBuybackResponse
	}{
		{
			name: "both prices are $1",
			postedAssetPairs: []common.AssetPair{
				common.GovStablePool,
				common.CollStablePool,
			},
			scenario: NeededCollScenario{
				protocolColl:    sdk.NewInt(700_000),
				priceCollStable: sdk.OneDec(),
				stableSupply:    sdk.NewInt(1_000_000),
				collRatio:       sdk.MustNewDecFromStr("0.6"),
				// neededUSD = (0.6 * 1000e3) - (700e3 *1) = -100_000
			},
			priceGovStable: sdk.OneDec(),
			accFunds: sdk.NewCoins(
				sdk.NewInt64Coin(common.GovDenom, 1_000_000_000),
			),
			moduleAccFunds: sdk.NewCoins(
				sdk.NewInt64Coin(common.CollDenom, 1_000_000_000),
			),

			expectedNeededUSD: sdk.NewDec(-100_000),
			msg: types.MsgBuyback{
				Creator: sample.AccAddress().String(),
				Gov:     sdk.NewCoin(common.GovDenom, sdk.NewInt(100_000)),
			},
			response: &types.MsgBuybackResponse{
				/*
					Coll.Amount = inUSD *  / priceCollStable
					  = 100_000 / priceCollStable
				*/
				Coll: sdk.NewCoin(common.CollDenom, sdk.NewInt(100_000)),
			},
			expectedPass: true,
		},
		{
			name: "arbitrary valid prices",
			postedAssetPairs: []common.AssetPair{
				common.GovStablePool,
				common.CollStablePool,
			},
			scenario: NeededCollScenario{
				protocolColl:    sdk.NewInt(850_000),
				priceCollStable: sdk.MustNewDecFromStr("1.099999"),
				stableSupply:    sdk.NewInt(1_000_000),
				collRatio:       sdk.MustNewDecFromStr("0.7"),
				// neededUSD =  (0.7 * 1000e3) - (850e3 *1.09999) = -234999.15
			},
			priceGovStable: sdk.NewDec(5),
			accFunds: sdk.NewCoins(
				sdk.NewInt64Coin(common.GovDenom, 1_000_000_000),
			),
			moduleAccFunds: sdk.NewCoins(
				sdk.NewInt64Coin(common.CollDenom, 1_000_000_000),
			),

			// Since 'neededUSD' is
			expectedNeededUSD: sdk.MustNewDecFromStr("-234999.15"),
			msg: types.MsgBuyback{
				Creator: sample.AccAddress().String(),
				Gov:     sdk.NewCoin(common.GovDenom, sdk.NewInt(50_000)),
			},
			response: &types.MsgBuybackResponse{
				/*
					Coll.Amount = inCollUSD  / priceCollStable
					  = msg.Gov.Amount * priceGovStable  / priceCollStable
					  = 50_000 * 5 / 1.099999
					  = 227272.93388448536 -> 227_272
				*/
				Coll: sdk.NewCoin(common.CollDenom, sdk.NewInt(227_272)),
			},
			expectedPass: true,
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			require.EqualValues(t, tc.expectedNeededUSD, tc.scenario.CalcNeededUSD())

			nibiruApp, ctx := testutil.NewNibiruApp(true)
			stablecoinKeeper := &nibiruApp.StablecoinKeeper
			require.NoError(t, stablecoinKeeper.SetCollRatio(ctx, tc.scenario.collRatio))
			require.NoError(t, nibiruApp.BankKeeper.MintCoins(
				ctx, types.ModuleName, sdk.NewCoins(
					sdk.NewCoin(common.CollDenom, tc.scenario.protocolColl),
					sdk.NewCoin(common.StableDenom, tc.scenario.stableSupply),
				),
			))
			// Fund module account
			err := nibiruApp.BankKeeper.MintCoins(ctx, types.ModuleName, tc.moduleAccFunds)
			if tc.expectedPass {
				require.NoError(t, err)
			}

			// Fund caller account
			caller, err := sdk.AccAddressFromBech32(tc.msg.Creator)
			if tc.expectedPass {
				require.NoError(t, err)
			}
			err = simapp.FundAccount(nibiruApp.BankKeeper, ctx, caller, tc.accFunds)
			if tc.expectedPass {
				require.NoError(t, err)
			}

			// Set up markets for the pricefeed keeper.
			oracle := sample.AccAddress()
			priceExpiry := ctx.BlockTime().Add(time.Hour)
			pricefeedParams := pricefeedTypes.Params{
				Pairs: []pricefeedTypes.Pair{
					{Token0: common.CollDenom, Token1: common.StableDenom,
						Oracles: []sdk.AccAddress{oracle}, Active: true},
					{Token0: common.GovDenom, Token1: common.StableDenom,
						Oracles: []sdk.AccAddress{oracle}, Active: true},
				}}
			nibiruApp.PriceKeeper.SetParams(ctx, pricefeedParams)

			prices := map[common.AssetPair]sdk.Dec{
				common.GovStablePool:  tc.priceGovStable,
				common.CollStablePool: tc.scenario.priceCollStable,
			}
			for _, pair := range tc.postedAssetPairs {
				_, err := nibiruApp.PriceKeeper.SetPrice(
					ctx, oracle, pair.Token0, pair.Token1, prices[pair], priceExpiry)
				require.NoError(t, err)

				// Update the 'CurrentPrice' posted by the oracles.
				err = nibiruApp.PriceKeeper.SetCurrentPrices(ctx, pair.Token0, pair.Token1)
				require.NoError(t, err, "Error posting price for pair: %d", pair.String())
			}

			// Post prices to each specified market with the oracle.
			for _, assetPair := range tc.postedAssetPairs {
				_, err := nibiruApp.PriceKeeper.SetPrice(
					ctx, oracle, assetPair.Token0, assetPair.Token1,
					prices[assetPair], priceExpiry)
				require.NoError(t, err)

				// Update the 'CurrentPrice' posted by the oracles.
				err = nibiruApp.PriceKeeper.SetCurrentPrices(
					ctx, assetPair.Token0, assetPair.Token1)
				require.NoError(
					t, err, "Error posting price for pair: %d", assetPair.String())
			}

			goCtx := sdk.WrapSDKContext(ctx)
			response, err := stablecoinKeeper.Buyback(goCtx, &tc.msg)
			if tc.expectedPass {
				require.NoError(t, err)
				require.EqualValues(t, tc.response, response)
			} else {
				require.Error(t, err)
			}
		},
		)
	}
}
