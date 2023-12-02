package keeper_test

import (
	"context"
	"github.com/alice/checkers/x/biggernum"
	"github.com/alice/checkers/x/biggernum/testutil"
	"github.com/stretchr/testify/require"
	"testing"

	keepertest "github.com/alice/checkers/testutil/keeper"
	"github.com/alice/checkers/x/biggernum/keeper"
	"github.com/alice/checkers/x/biggernum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BiggernumKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func setupMsgServerWithInitGenesis(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.BiggernumKeeper(t)
	biggernum.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestGame(t *testing.T) {
	msgServer, keeper, ctx := setupMsgServerWithInitGenesis(t)
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	createGameResp, err := msgServer.CreateGame(ctx, &types.MsgCreateGame{
		Creator: testutil.Alice,
		Player1: testutil.Alice,
		Player2: testutil.Bob,
	})
	require.NoError(t, err)
	require.Equal(t, "0", createGameResp.GameId)

	gameList := keeper.GetAllGames(sdkCtx)
	require.Len(t, gameList, 1)
	game := gameList[0]
	require.Equal(t, "0", game.Index)
	require.Equal(t, testutil.Alice, game.Player1)
	require.Equal(t, testutil.Bob, game.Player2)
	require.Equal(t, uint64(0), game.Move1)
	require.Equal(t, uint64(0), game.Move2)
	require.Equal(t, uint64(1), game.PlayerToMove)
	require.Equal(t, "", game.Winner)

	_, err = msgServer.PlayMove(ctx, &types.MsgPlayMove{
		Creator: testutil.Alice,
		GameId:  "0",
		Number:  42,
	})
	require.NoError(t, err)

	gameList = keeper.GetAllGames(sdkCtx)
	require.Len(t, gameList, 1)
	game = gameList[0]
	//gameInBytes, err := json.Marshal(game)
	//require.NoError(t, err)
	//fmt.Printf("gameInBytes: %s\n", gameInBytes)
	require.Equal(t, "0", game.Index)
	require.Equal(t, testutil.Alice, game.Player1)
	require.Equal(t, testutil.Bob, game.Player2)
	require.Equal(t, uint64(42), game.Move1)
	require.Equal(t, uint64(0), game.Move2)
	require.Equal(t, uint64(2), game.PlayerToMove)
	require.Equal(t, "", game.Winner)

	_, err = msgServer.PlayMove(ctx, &types.MsgPlayMove{
		Creator: testutil.Bob,
		GameId:  "0",
		Number:  43,
	})
	require.NoError(t, err)

	gameList = keeper.GetAllGames(sdkCtx)
	require.Len(t, gameList, 1)
	game = gameList[0]
	require.Equal(t, "0", game.Index)
	require.Equal(t, testutil.Alice, game.Player1)
	require.Equal(t, testutil.Bob, game.Player2)
	require.Equal(t, uint64(42), game.Move1)
	require.Equal(t, uint64(43), game.Move2)
	require.Equal(t, uint64(3), game.PlayerToMove)
	require.Equal(t, testutil.Bob, game.Winner)
}
