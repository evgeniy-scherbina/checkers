package keeper

import (
	"github.com/alice/checkers/x/biggernum/types"
)

var _ types.QueryServer = Keeper{}
