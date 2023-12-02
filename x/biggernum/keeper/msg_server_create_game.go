package keeper

import (
	"context"
	"strconv"

	"github.com/alice/checkers/x/biggernum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	// get game index
	systemInfo, found := k.GetSystemInfo(ctx)
	if !found {
		panic("system info is not found")
	}
	newIndex := strconv.FormatUint(systemInfo.NextId, 10)

	// store the game
	k.SetGames(ctx, types.Games{
		Index:        newIndex,
		Player1:      msg.Player1,
		Player2:      msg.Player2,
		PlayerToMove: 1,
	})

	// update index
	systemInfo.NextId++
	k.SetSystemInfo(ctx, systemInfo)

	return &types.MsgCreateGameResponse{
		GameId: newIndex,
	}, nil
}
