package keeper

import (
	"github.com/alice/checkers/x/newcheckers/types"
)

var _ types.QueryServer = Keeper{}
