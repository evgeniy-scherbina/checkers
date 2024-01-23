package keeper

import (
	"context"
	"github.com/alice/checkers/x/oddnum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Transfer(goCtx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	// ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins

	fromAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}
	toAddr, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return nil, err
	}
	coin := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(msg.Amount)))
	coins := sdk.NewCoins(coin)

	err = k.bank.SendCoins(ctx, fromAddr, toAddr, coins)
	if err != nil {
		return nil, err
	}

	return &types.MsgTransferResponse{}, nil
}
