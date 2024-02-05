package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/alice/checkers/testutil/keeper"
	"github.com/alice/checkers/testutil/nullify"
	"github.com/alice/checkers/x/monobank/keeper"
	"github.com/alice/checkers/x/monobank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNBalance(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Balance {
	items := make([]types.Balance, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetBalance(ctx, items[i])
	}
	return items
}

func TestBalanceGet(t *testing.T) {
	keeper, ctx := keepertest.MonobankKeeper(t)
	items := createNBalance(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetBalance(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestBalanceRemove(t *testing.T) {
	keeper, ctx := keepertest.MonobankKeeper(t)
	items := createNBalance(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBalance(ctx,
			item.Address,
		)
		_, found := keeper.GetBalance(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestBalanceGetAll(t *testing.T) {
	keeper, ctx := keepertest.MonobankKeeper(t)
	items := createNBalance(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllBalance(ctx)),
	)
}
