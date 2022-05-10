package keeper

import (
	"fmt"

	"github.com/NibiruChain/nibiru/x/common"
	"github.com/NibiruChain/nibiru/x/vpool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// addReserveSnapshot adds a snapshot of the current pool status and blocktime and blocknum.
func (k Keeper) addReserveSnapshot(ctx sdk.Context, pool *types.Pool) error {
	blockNumber := ctx.BlockHeight()
	lastSnapshot, lastCounter, err := k.getLastReserveSnapshot(ctx, common.TokenPair(pool.Pair))
	if err != nil {
		return err
	}

	if blockNumber == lastSnapshot.BlockNumber {
		err = k.updateSnapshot(ctx, lastCounter, pool)
		if err != nil {
			return fmt.Errorf("error updating snapshot: %w", err)
		}
	} else {
		err = k.saveReserveSnapshot(ctx, lastCounter, pool)
		if err != nil {
			return fmt.Errorf("error saving snapshot: %w", err)
		}
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventSnapshotSaved,
			sdk.NewAttribute(types.AttributeBlockHeight, fmt.Sprintf("%d", ctx.BlockHeight())),
			sdk.NewAttribute(types.AttributeQuoteReserve, pool.QuoteAssetReserve.String()),
			sdk.NewAttribute(types.AttributeBaseReserve, pool.BaseAssetReserve.String()),
		),
	)

	return nil
}

// saveReserveSnapshot saves reserve snapshot and increments counter
func (k Keeper) saveReserveSnapshot(ctx sdk.Context, lastCounter int64, pool *types.Pool) error {
	counter := lastCounter + 1

	err := k.saveSnapshotInStore(ctx, pool, counter)
	if err != nil {
		return err
	}

	k.updateSnapshotCounter(ctx, common.TokenPair(pool.Pair), counter)

	return nil
}

// updateSnapshot saves the snapshot but does not increase the counter
func (k Keeper) updateSnapshot(ctx sdk.Context, counter int64, pool *types.Pool) error {
	err := k.saveSnapshotInStore(ctx, pool, counter)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) saveSnapshotInStore(ctx sdk.Context, pool *types.Pool, counter int64) error {
	snapshot := &types.ReserveSnapshot{
		BaseAssetReserve:  pool.BaseAssetReserve,
		QuoteAssetReserve: pool.QuoteAssetReserve,
		Timestamp:         ctx.BlockTime().Unix(),
		BlockNumber:       ctx.BlockHeight(),
	}
	bz, err := k.codec.Marshal(snapshot)
	if err != nil {
		return err
	}

	store := k.getStore(ctx)
	store.Set(types.GetPoolReserveSnapshotKey(counter), bz)

	return nil
}

// getSnapshotCounter returns the counter and if it has been found or not.
func (k Keeper) getSnapshotCounter(ctx sdk.Context, pair common.TokenPair) (int64, bool) {
	store := k.getStore(ctx)

	bz := store.
		Get(types.GetPoolReserveSnapshotCounter(pair))
	if bz == nil {
		return 0, false
	}

	sc := sdk.BigEndianToUint64(bz)

	return int64(sc), true
}

func (k Keeper) updateSnapshotCounter(ctx sdk.Context, pair common.TokenPair, counter int64) {
	store := k.getStore(ctx)

	store.Set(types.GetPoolReserveSnapshotCounter(pair), sdk.Uint64ToBigEndian(uint64(counter)))
}

// getLastReserveSnapshot returns the last snapshot that was saved
func (k Keeper) getLastReserveSnapshot(ctx sdk.Context, pair common.TokenPair) (types.ReserveSnapshot, int64, error) {
	counter, found := k.getSnapshotCounter(ctx, pair)
	if !found {
		return types.ReserveSnapshot{}, counter, types.ErrNoLastSnapshotSaved
	}

	snapshot, err := k.getSnapshotByCounter(ctx, counter)
	if err != nil {
		return types.ReserveSnapshot{}, counter, types.ErrNoLastSnapshotSaved
	}

	return snapshot, counter, err
}

// getSnapshotByCounter returns the snapshot saved by counter num
func (k Keeper) getSnapshotByCounter(ctx sdk.Context, counter int64) (types.ReserveSnapshot, error) {
	store := k.getStore(ctx)
	bz := store.Get(types.GetPoolReserveSnapshotKey(counter))
	if bz == nil {
		return types.ReserveSnapshot{}, types.ErrNoLastSnapshotSaved.
			Wrap(fmt.Sprintf("snapshot with counter %d was not found", counter))
	}

	var snapshot types.ReserveSnapshot
	err := k.codec.Unmarshal(bz, &snapshot)
	if err != nil {
		return types.ReserveSnapshot{}, fmt.Errorf("problem decoding snapshot")
	}

	return snapshot, nil
}
