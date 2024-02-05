package keeper

import (
	"github.com/alice/checkers/x/monobank/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetBalance set a specific balance in the store from its index
func (k Keeper) SetBalance(ctx sdk.Context, balance types.Balance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BalanceKeyPrefix))
	b := k.cdc.MustMarshal(&balance)
	store.Set(types.BalanceKey(
		balance.Address,
	), b)
}

// GetBalance returns a balance from its index
func (k Keeper) GetBalance(
	ctx sdk.Context,
	address string,

) (val types.Balance, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BalanceKeyPrefix))

	b := store.Get(types.BalanceKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBalance removes a balance from the store
func (k Keeper) RemoveBalance(
	ctx sdk.Context,
	address string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BalanceKeyPrefix))
	store.Delete(types.BalanceKey(
		address,
	))
}

// GetAllBalance returns all balance
func (k Keeper) GetAllBalance(ctx sdk.Context) (list []types.Balance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BalanceKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Balance
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
