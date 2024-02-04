package keeper

import (
	"github.com/alice/checkers/x/monobank/types"
)

var _ types.QueryServer = Keeper{}
