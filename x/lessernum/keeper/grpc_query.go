package keeper

import (
	"github.com/alice/checkers/x/lessernum/types"
)

var _ types.QueryServer = Keeper{}
