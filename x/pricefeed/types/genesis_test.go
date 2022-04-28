package types

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmtypes "github.com/tendermint/tendermint/types"
)

func TestGenesisState_Validate(t *testing.T) {
	mockPrivKey := tmtypes.NewMockPV()
	pubkey, err := mockPrivKey.GetPubKey()
	require.NoError(t, err)
	addr := sdk.AccAddress(pubkey.Address())
	now := time.Now()

	for _, tc := range []struct {
		desc     string
		genState *GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: DefaultGenesis(),
			valid:    true,
		},
		{
			desc:     "valid genesis state",
			genState: &GenesisState{},
			valid:    true,
		},
		{
			desc: "dup market param",
			genState: NewGenesisState(
				NewParams([]Pair{
					{"xrp", "bnb", []sdk.AccAddress{addr}, true},
					{"xrp", "bnb", []sdk.AccAddress{addr}, true},
				}),
				[]PostedPrice{NewPostedPrice("xrp", addr, sdk.OneDec(), now)},
			),
			valid: false,
		},
		{
			desc: "invalid posted price",
			genState: NewGenesisState(
				NewParams([]Pair{}),
				[]PostedPrice{NewPostedPrice("xrp", nil, sdk.OneDec(), now)},
			),
			valid: false,
		},
		{
			desc: "duplicated posted price",
			genState: NewGenesisState(
				NewParams([]Pair{}),
				[]PostedPrice{
					NewPostedPrice("xrp", addr, sdk.OneDec(), now),
					NewPostedPrice("xrp", addr, sdk.OneDec(), now),
				},
			),
			valid: false,
		},

		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
