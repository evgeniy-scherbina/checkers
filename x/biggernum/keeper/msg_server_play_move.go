package keeper

import (
	"context"
	"fmt"
	"github.com/alice/checkers/x/biggernum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PlayMove(goCtx context.Context, msg *types.MsgPlayMove) (*types.MsgPlayMoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	// get game from storage
	game, found := k.GetGames(ctx, msg.GetGameId())
	if !found {
		return nil, fmt.Errorf("invalid id")
	}

	// check if move is done by correct player
	if game.PlayerToMove == 1 {
		if msg.Creator != game.Player1 {
			return nil, fmt.Errorf("not your turn")
		}
	} else if game.PlayerToMove == 2 {
		if msg.Creator != game.Player2 {
			return nil, fmt.Errorf("not your turn")
		}
	} else {
		panic("player to move should be either 1 or 2")
	}

	// process move
	if game.PlayerToMove == 1 {
		game.Move1 = msg.Number
		game.PlayerToMove++
	} else if game.PlayerToMove == 2 {
		game.Move2 = msg.Number
	}

	// process finished game
	if game.PlayerToMove == 2 {
		var winner string
		if game.Move1 >= game.Move2 {
			winner = game.Player1
		} else {
			winner = game.Player2
		}

		game.Winner = winner
	}

	k.SetGames(ctx, game)

	return &types.MsgPlayMoveResponse{}, nil
}
