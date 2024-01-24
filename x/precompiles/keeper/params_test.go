package keeper_test

import (
	"testing"

	testkeeper "github.com/alice/checkers/testutil/keeper"
	"github.com/alice/checkers/x/precompiles/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PrecompilesKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
