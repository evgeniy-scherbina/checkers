package keeper

import (
	"github.com/alice/checkers/x/sum/types"
)

var _ types.QueryServer = Keeper{}
