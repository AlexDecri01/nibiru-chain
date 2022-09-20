package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"

	"github.com/NibiruChain/nibiru/x/common"
)

func NewReserveSnapshot(
	pair common.AssetPair,
	baseAssetReserve, quoteAssetReserve sdk.Dec,
	blockTime time.Time,
	blockHeight int64,
) ReserveSnapshot {
	return ReserveSnapshot{
		Pair:              pair,
		BaseAssetReserve:  baseAssetReserve,
		QuoteAssetReserve: quoteAssetReserve,
		TimestampMs:       blockTime.UnixMilli(),
		BlockNumber:       blockHeight,
	}
}
