package keeper

import (
	"context"

	"github.com/alice/checkers/x/oddnum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) PlayMove(goCtx context.Context, msg *types.MsgPlayMove) (*types.MsgPlayMoveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	storedGame, found := k.Keeper.GetStoredGame(ctx, msg.GameIndex)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrGameNotFound, "%s", msg.GameIndex)
	}
	_ = storedGame

	black, err := storedGame.GetBlackAddress()
	if err != nil {
		panic(err.Error())
	}
	err = k.bank.SendCoinsFromAccountToModule(ctx, black, types.ModuleName, sdk.NewCoins(storedGame.GetWagerCoin()))
	if err != nil {
		return nil, sdkerrors.Wrapf(err, types.ErrBlackCannotPay.Error())
	}

	return &types.MsgPlayMoveResponse{}, nil
}
