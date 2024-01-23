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

	ErrBlackCannotPay    = sdkerrors.Register(ModuleName, 1110, "black cannot pay the wager")
	ErrRedCannotPay      = sdkerrors.Register(ModuleName, 1111, "red cannot pay the wager")
	ErrNothingToPay      = sdkerrors.Register(ModuleName, 1112, "there is nothing to pay, should not have been called")
	ErrCannotRefundWager = sdkerrors.Register(ModuleName, 1113, "cannot refund wager to: %s")
	ErrCannotPayWinnings = sdkerrors.Register(ModuleName, 1114, "cannot pay winnings to winner: %s")
	ErrNotInRefundState  = sdkerrors.Register(ModuleName, 1115, "game is not in a state to refund, move count: %d")
)
