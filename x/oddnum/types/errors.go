package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/oddnum module sentinel errors
var (
	ErrInvalidBlack = sdkerrors.Register(ModuleName, 1100, "black address is invalid: %s")
	ErrInvalidRed   = sdkerrors.Register(ModuleName, 1101, "red address is invalid: %s")

	ErrGameNotFound     = sdkerrors.Register(ModuleName, 1106, "game by id not found")
	ErrCreatorNotPlayer = sdkerrors.Register(ModuleName, 1107, "message creator is not a player")
	ErrNotPlayerTurn    = sdkerrors.Register(ModuleName, 1108, "player tried to play out of turn")
)
