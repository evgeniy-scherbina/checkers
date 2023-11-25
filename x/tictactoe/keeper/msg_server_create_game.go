package keeper

import (
	"context"
	"github.com/alice/checkers/x/tictactoe/rules"
	"strconv"

	"github.com/alice/checkers/x/tictactoe/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	// 2. Get the new game's ID with the Keeper.GetSystemInfo function
	// created by the ignite scaffold single systemInfo... command:
	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}
	newIndex := strconv.FormatUint(systemInfo.NextId, 10)

	// 3. Create the object to be stored:
	newGame := rules.NewGame()
	storedGame := types.StoredGame{
		Index:    newIndex,
		Board:    string(newGame.MustSerialize()),
		NextTurn: string(newGame.NextTurn()),
		XPlayer:  msg.XPlayer,
		OPlayer:  msg.OPlayer,
	}

	// 4. Confirm that the values in the object are correct by checking the validity of the players' addresses:
	err := storedGame.Validate()
	if err != nil {
		return nil, err
	}

	// 5. Save the StoredGame object using the Keeper.SetStoredGame function created by the ignite scaffold map storedGame... command:
	k.Keeper.SetStoredGame(ctx, storedGame)

	// 6. Prepare the ground for the next game using the Keeper.SetSystemInfo function created by Ignite CLI:
	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)

	// 7. Return the newly created ID for reference:
	return &types.MsgCreateGameResponse{
		GameIndex: newIndex,
	}, nil
}
