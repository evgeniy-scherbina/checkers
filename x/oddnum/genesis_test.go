package oddnum_test

import (
	"testing"

	keepertest "github.com/alice/checkers/testutil/keeper"
	"github.com/alice/checkers/testutil/nullify"
	"github.com/alice/checkers/x/oddnum"
	"github.com/alice/checkers/x/oddnum/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SystemInfo: &types.SystemInfo{
			NextId: 2,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OddnumKeeper(t)
	oddnum.InitGenesis(ctx, *k, genesisState)
	got := oddnum.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	// this line is used by starport scaffolding # genesis/test/assert
}
