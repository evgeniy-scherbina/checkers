package keeper

import (
	"github.com/alice/checkers/x/tictactoe/types"
)

var _ types.QueryServer = Keeper{}
