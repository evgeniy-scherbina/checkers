package keeper

import (
	"context"

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

	return &types.MsgPlayMoveResponse{}, nil
}

func isPlayerTurn(storedGame *types.StoredGame, player string) bool {
	switch storedGame. {

	}
}