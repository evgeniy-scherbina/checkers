package keeper

import (
	"context"
	"strconv"

	"github.com/alice/checkers/x/lessernum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	systemInfo, found := k.GetSystemInfo(ctx)
	if !found {
		panic("can't get system info")
	}
	nextId := strconv.FormatUint(systemInfo.NextId, 10)

	game := types.StoredGame{
		Index:        nextId,
		Player1:      msg.Player1,
		Player2:      msg.Player2,
		PlayerToMove: 1,
		Wager:        msg.Wager,
	}
	k.SetStoredGame(ctx, game)

	systemInfo.NextId++
	k.SetSystemInfo(ctx, systemInfo)

	return &types.MsgCreateGameResponse{
		GameId: nextId,
	}, nil
}
