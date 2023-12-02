package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/lessernum module sentinel errors
var (
	ErrGameNotFound = sdkerrors.Register(ModuleName, 1100, "game not found")
)
