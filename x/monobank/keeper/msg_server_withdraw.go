package keeper

import (
	"context"
	"errors"
	"github.com/alice/checkers/x/monobank/testutil"
	"github.com/alice/checkers/x/monobank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	withdrawer, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	balance, found := k.GetBalance(ctx, withdrawer.String())
	if !found || balance.Balance < msg.Amount {
		return nil, errors.New("insufficient funds")
	}

	balance.Balance -= msg.Amount
	k.SetBalance(ctx, balance)

	err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, withdrawer, testutil.CoinsOf(msg.Amount))
	if err != nil {
		return nil, err
	}

	return &types.MsgWithdrawResponse{}, nil
}
