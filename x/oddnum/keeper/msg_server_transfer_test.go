package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/alice/checkers/x/oddnum/testutil"
	"github.com/alice/checkers/x/oddnum/types"
)

func TestTransfer(t *testing.T) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankEscrowKeeper(ctrl)

	msgServer, goCtx := setupMsgServerWithBankKeeper(t, bankMock)

	bankMock.EXPECT().
		SendCoins(sdk.UnwrapSDKContext(goCtx), gomock.Any(), gomock.Any(), gomock.Any()).
		Times(1)

	alice := testutil.Alice
	bob := testutil.Bob

	_, err := msgServer.Transfer(goCtx, &types.MsgTransfer{
		Creator: alice,
		From:    alice,
		To:      bob,
		Amount:  100_000,
	})
	require.NoError(t, err)
}
