package keeper

import (
	"github.com/alice/checkers/x/precompiles/types"
)

var _ types.QueryServer = Keeper{}
