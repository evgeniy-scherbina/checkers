package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/alice/checkers/testutil/keeper"
	"github.com/alice/checkers/x/oddnum/keeper"
	"github.com/alice/checkers/x/oddnum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.OddnumKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func setupMsgServerWithBankKeeper(t testing.TB, bank types.BankEscrowKeeper) (types.MsgServer, context.Context) {
	k, ctx := keepertest.OddnumKeeperWithBankKeeper(t, bank)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
