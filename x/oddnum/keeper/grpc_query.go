package keeper

import (
	"github.com/alice/checkers/x/oddnum/types"
)

var _ types.QueryServer = Keeper{}
