package keeper_test

import (
	keepertest "github.com/alice/checkers/testutil/keeper"
	"github.com/alice/checkers/x/monobank/keeper"
	"github.com/alice/checkers/x/monobank/testutil"
	"github.com/alice/checkers/x/monobank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDepositAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankEscrowKeeper(ctrl)

	k, ctx := keepertest.MonobankKeeperWithMocks(t, bankMock)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	aliceAddr, err := sdk.AccAddressFromBech32(testutil.Alice)
	require.NoError(t, err)

	bankMock.EXPECT().
		SendCoinsFromAccountToModule(ctx, aliceAddr, types.ModuleName, testutil.CoinsOf(1)).
		AnyTimes()

	_, err = msgServer.Deposit(goCtx, &types.MsgDeposit{
		Creator: testutil.Alice,
		Amount:  1,
	})
	require.NoError(t, err)
}
