package keeper

import (
	"context"
	"fmt"
	"github.com/alice/checkers/x/checkers/rules"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ForfeitExpiredGames(goCtx context.Context) {
	// TODO

	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1. Prepare useful information:
	opponents := map[string]string{
		rules.PieceStrings[rules.BLACK_PLAYER]: rules.PieceStrings[rules.RED_PLAYER],
		rules.PieceStrings[rules.RED_PLAYER]:   rules.PieceStrings[rules.BLACK_PLAYER],
	}

	// 2. Initialize the parameters before entering the loop:
	systemInfo, found := k.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}

	gameIndex := systemInfo.FifoHeadIndex
	var storedGame types.StoredGame

	// 3. Enter the loop:
	for {
		// 1. Start with a loop breaking condition, if your cursor has reached the end of the FIFO:
		if gameIndex == types.NoFifoIndex {
			break
		}

		// 2. Fetch the expired game candidate and its deadline:
		storedGame, found = k.GetStoredGame(ctx, gameIndex)
		if !found {
			panic("Fifo head game not found " + systemInfo.FifoHeadIndex)
		}
		deadline, err := storedGame.GetDeadlineAsTime()
		if err != nil {
			panic(err)
		}

		// 3. Test for expiration:
		if deadline.Before(ctx.BlockTime()) {
			// 1. If the game has expired, remove it from the FIFO:
			k.RemoveFromFifo(ctx, &storedGame, &systemInfo)

			// 2. Check whether the game is worth keeping. If it is, set the winner as the opponent of the player whose
			// turn it is, remove the board, and save:
			lastBoard := storedGame.Board
			if storedGame.MoveCount <= 1 {
				// No point in keeping a game that was never really played
				k.RemoveStoredGame(ctx, gameIndex)
				if storedGame.MoveCount == 1 {
					k.MustRefundWager(ctx, &storedGame)
				}
			} else {
				storedGame.Winner, found = opponents[storedGame.Turn]
				if !found {
					panic(fmt.Sprintf(types.ErrCannotFindWinnerByColor.Error(), storedGame.Turn))
				}
				k.MustPayWinnings(ctx, &storedGame)
				storedGame.Board = ""
				k.SetStoredGame(ctx, storedGame)
			}

			// 3. Emit the relevant event:
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(types.GameForfeitedEventType,
					sdk.NewAttribute(types.GameForfeitedEventGameIndex, gameIndex),
					sdk.NewAttribute(types.GameForfeitedEventWinner, storedGame.Winner),
					sdk.NewAttribute(types.GameForfeitedEventBoard, lastBoard),
				),
			)

			// 4. Move along the FIFO for the next run of the loop:
			gameIndex = systemInfo.FifoHeadIndex
		} else {
			// All other games after are active anyway
			break
		}
	}

	k.SetSystemInfo(ctx, systemInfo)
}
