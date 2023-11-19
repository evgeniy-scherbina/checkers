package keeper

import (
	"context"
	"github.com/alice/checkers/x/checkers/rules"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PlayMove(goCtx context.Context, msg *types.MsgPlayMove) (*types.MsgPlayMoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	// 1. Fetch the stored game information using the Keeper.GetStoredGame function created by Ignite CLI:
	storedGame, found := k.Keeper.GetStoredGame(ctx, msg.GameIndex)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "%s", msg.GameIndex)
	}

	if storedGame.Winner != rules.PieceStrings[rules.NO_PLAYER] {
		return nil, types.ErrGameFinished
	}

	// 2. Is the player legitimate? Check with:
	isBlack := storedGame.Black == msg.Creator
	isRed := storedGame.Red == msg.Creator
	var player rules.Player
	if !isBlack && !isRed {
		return nil, sdkerrors.Wrapf(types.ErrCreatorNotPlayer, "%s", msg.Creator)
	} else if isBlack && isRed {
		player = rules.StringPieces[storedGame.Turn].Player
	} else if isBlack {
		player = rules.BLACK_PLAYER
	} else {
		player = rules.RED_PLAYER
	}

	// 3. Instantiate the board in order to implement the rules:
	game, err := storedGame.ParseGame()
	if err != nil {
		panic(err.Error())
	}

	// 4. Is it the player's turn? Check using the rules file's own TurnIs function:
	if !game.TurnIs(player) {
		return nil, sdkerrors.Wrapf(types.ErrNotPlayerTurn, "%s", player)
	}

	// 5. Properly conduct the move, using the rules' Move function:
	captured, moveErr := game.Move(
		rules.Pos{
			X: int(msg.FromX),
			Y: int(msg.FromY),
		},
		rules.Pos{
			X: int(msg.ToX),
			Y: int(msg.ToY),
		},
	)
	if moveErr != nil {
		return nil, sdkerrors.Wrapf(types.ErrWrongMove, moveErr.Error())
	}

	// 6. Prepare the updated board to be stored and store the information:
	storedGame.Winner = rules.PieceStrings[game.Winner()]

	lastBoard := game.String()
	if storedGame.Winner == rules.PieceStrings[rules.NO_PLAYER] {
		storedGame.Board = lastBoard
	} else {
		storedGame.Board = ""
	}

	storedGame.Deadline = types.FormatDeadline(types.GetNextDeadline(ctx))
	storedGame.Turn = rules.PieceStrings[game.Turn]
	k.Keeper.SetStoredGame(ctx, storedGame)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.MovePlayedEventType,
			sdk.NewAttribute(types.MovePlayedEventCreator, msg.Creator),
			sdk.NewAttribute(types.MovePlayedEventGameIndex, msg.GameIndex),
			sdk.NewAttribute(types.MovePlayedEventCapturedX, strconv.FormatInt(int64(captured.X), 10)),
			sdk.NewAttribute(types.MovePlayedEventCapturedY, strconv.FormatInt(int64(captured.Y), 10)),
			sdk.NewAttribute(types.MovePlayedEventWinner, rules.PieceStrings[game.Winner()]),
			sdk.NewAttribute(types.MovePlayedEventBoard, lastBoard),
		),
	)

	// 7. Return relevant information regarding the move's result:
	return &types.MsgPlayMoveResponse{
		CapturedX: int32(captured.X),
		CapturedY: int32(captured.Y),
		Winner:    rules.PieceStrings[game.Winner()],
	}, nil
}
