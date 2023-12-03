package keeper

import (
	"context"
	"fmt"
	"github.com/alice/checkers/x/lessernum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PlayMove(goCtx context.Context, msg *types.MsgPlayMove) (*types.MsgPlayMoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	storedGame, found := k.GetStoredGame(ctx, msg.GameId)
	if !found {
		return nil, types.ErrGameNotFound
	}

	active := isGameActive(&storedGame)
	if !active {
		return nil, fmt.Errorf("game is not active")
	}

	ok := isPlayerTurn(&storedGame, msg.Creator)
	if !ok {
		return nil, fmt.Errorf("not your turn")
	}

	makeMove(&storedGame, msg.Number)
	k.SetStoredGame(ctx, storedGame)

	return &types.MsgPlayMoveResponse{}, nil
}

func isGameActive(storedGame *types.StoredGame) bool {
	return storedGame.PlayerToMove == 1 || storedGame.PlayerToMove == 2
}

func isPlayerTurn(storedGame *types.StoredGame, player string) bool {
	switch storedGame.PlayerToMove {
	case 1:
		return storedGame.Player1 == player
	case 2:
		return storedGame.Player2 == player
	default:
		return false
	}
}

func makeMove(storedGame *types.StoredGame, number uint64) {
	switch storedGame.PlayerToMove {
	case 1:
		storedGame.Move1 = number
		storedGame.PlayerToMove++
	case 2:
		storedGame.Move2 = number

		if storedGame.Move1 < storedGame.Move2 {
			storedGame.Winner = storedGame.Player1
		} else {
			storedGame.Winner = storedGame.Player2
		}

		storedGame.PlayerToMove++
	}
}
