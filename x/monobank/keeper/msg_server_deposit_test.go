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
	bankMock.EXPECT().
		SendCoinsFromAccountToModule(ctx, aliceAddr, types.ModuleName, testutil.CoinsOf(2)).
		AnyTimes()

	bankMock.EXPECT().
		SendCoinsFromModuleToAccount(ctx, types.ModuleName, aliceAddr, testutil.CoinsOf(1)).
		AnyTimes()
	bankMock.EXPECT().
		SendCoinsFromModuleToAccount(ctx, types.ModuleName, aliceAddr, testutil.CoinsOf(2)).
		AnyTimes()

	// Deposit 1 token
	{
		_, err = msgServer.Deposit(goCtx, &types.MsgDeposit{
			Creator: testutil.Alice,
			Amount:  1,
		})
		require.NoError(t, err)

		balance, found := k.GetBalance(ctx, testutil.Alice)
		require.True(t, found)
		require.Equal(t, uint64(1), balance.Balance)
	}

	// Deposit 2 more tokens
	{
		_, err = msgServer.Deposit(goCtx, &types.MsgDeposit{
			Creator: testutil.Alice,
			Amount:  2,
		})
		require.NoError(t, err)

		balance, found := k.GetBalance(ctx, testutil.Alice)
		require.True(t, found)
		require.Equal(t, uint64(3), balance.Balance)
	}

	// Withdraw 1 token
	{
		_, err = msgServer.Withdraw(goCtx, &types.MsgWithdraw{
			Creator: testutil.Alice,
			Amount:  1,
		})
		require.NoError(t, err)

		balance, found := k.GetBalance(ctx, testutil.Alice)
		require.True(t, found)
		require.Equal(t, uint64(2), balance.Balance)
	}

	// Withdraw 2 more tokens
	{
		_, err = msgServer.Withdraw(goCtx, &types.MsgWithdraw{
			Creator: testutil.Alice,
			Amount:  2,
		})
		require.NoError(t, err)

		balance, found := k.GetBalance(ctx, testutil.Alice)
		require.True(t, found)
		require.Equal(t, uint64(0), balance.Balance)
	}
}
