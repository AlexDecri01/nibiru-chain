package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	tollRatio sdk.Dec,
	spreadRatio sdk.Dec,
) Params {
	million := sdk.NewDec(1_000_000)

	tollRatioInt := tollRatio.Mul(million).RoundInt().Int64()
	spreadRationInt := spreadRatio.Mul(million).RoundInt().Int64()

	return Params{
		TollRatio:   tollRatioInt,
		SpreadRatio: spreadRationInt,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	tollRatio := sdk.MustNewDecFromStr("0.001")
	spreadRatio := sdk.MustNewDecFromStr("0.001")

	return NewParams(tollRatio, spreadRatio)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(
			[]byte("TollRatio"),
			&p.TollRatio,
			validateTollRatio,
		),
		paramtypes.NewParamSetPair(
			[]byte("SpreadRatio"),
			&p.SpreadRatio,
			validateSpreadRatio,
		),
	}
}

// Validate validates the set of params
func (p *Params) Validate() error {
	err := validateTollRatio(p.TollRatio)
	if err != nil {
		return err
	}

	return validateSpreadRatio(p.SpreadRatio)
}

func (p *Params) GetSpreadRatioAsDec() sdk.Dec {
	return sdk.NewIntFromUint64(uint64(p.SpreadRatio)).
		ToDec().Quo(sdk.MustNewDecFromStr("1000000"))
}

func (p *Params) GetTollRatioAsDec() sdk.Dec {
	return sdk.NewIntFromUint64(uint64(p.TollRatio)).
		ToDec().Quo(sdk.MustNewDecFromStr("1000000"))
}

func validateTollRatio(i interface{}) error {
	tollRatio, err := getAsInt64(i)
	if err != nil {
		return err
	}

	if tollRatio > 1_000_000 {
		return fmt.Errorf("Toll ratio is above max value(1e6): %d", tollRatio)
	} else if tollRatio < 0 {
		return fmt.Errorf("Toll Ratio is negative: %d", tollRatio)
	} else {
		return nil
	}
}

func validateSpreadRatio(i interface{}) error {
	spreadRatio, err := getAsInt64(i)
	if err != nil {
		return err
	}

	if spreadRatio > 1_000_000 {
		return fmt.Errorf("spread ratio is above max value(1e6): %d", spreadRatio)
	} else if spreadRatio < 0 {
		return fmt.Errorf("spread ratio is negative: %d", spreadRatio)
	} else {
		return nil
	}
}

func getAsInt64(i interface{}) (int64, error) {
	value, ok := i.(int64)
	if !ok {
		return 0, fmt.Errorf("invalid parameter type: %T", i)
	}
	return value, nil
}
