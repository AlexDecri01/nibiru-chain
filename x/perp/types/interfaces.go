package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ----------------------------------------------------------
// ClearingHouse Interface
// ----------------------------------------------------------

type IClearingHouse interface {
	ClearPosition(ctx sdk.Context, vpool IVirtualPool, owner string) error
	GetPosition(
		ctx sdk.Context, vpool IVirtualPool, owner string,
	) (*Position, error)
	SetPosition(
		ctx sdk.Context, vpool IVirtualPool, owner string, position *Position,
	)
}

// ----------------------------------------------------------
// Vpool Interface
// ----------------------------------------------------------

type VirtualPoolDirection uint8

const (
	VirtualPoolDirection_AddToAMM = iota
	VirtualPoolDirection_RemoveFromAMM
)

type IVirtualPool interface {
	Pair() string
	QuoteTokenDenom() string
	SwapInput(
		ctx sdk.Context, ammDir VirtualPoolDirection, inputAmount,
		minOutputAmount sdk.Int, canOverFluctuationLimit bool,
	) (sdk.Int, error)
	SwapOutput(
		ctx sdk.Context, dir VirtualPoolDirection, abs sdk.Int, limit sdk.Int,
	) (sdk.Int, error)
	GetOpenInterestNotionalCap(ctx sdk.Context) (sdk.Int, error)
	GetMaxHoldingBaseAsset(ctx sdk.Context) (sdk.Int, error)
	GetOutputTWAP(ctx sdk.Context, dir VirtualPoolDirection, abs sdk.Int,
	) (sdk.Int, error)
	GetOutputPrice(ctx sdk.Context, dir VirtualPoolDirection, abs sdk.Int,
	) (sdk.Int, error)
	GetUnderlyingPrice(ctx sdk.Context) (sdk.Dec, error)
	GetSpotPrice(ctx sdk.Context) (sdk.Int, error)
	CalcFee(ctx sdk.Context, quoteAmt sdk.Int) (toll sdk.Int, spread sdk.Int, err error)
}
