package types_test

import (
	"testing"

	"github.com/alice/checkers/x/tictactoe/rules"
	"github.com/alice/checkers/x/tictactoe/testutil"
	"github.com/alice/checkers/x/tictactoe/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = testutil.Alice
	bob   = testutil.Bob
)

func GetStoredGame1() *types.StoredGame {
	g := rules.NewGame()
	serializedBoard, err := g.Serialize()
	if err != nil {
		panic(err)
	}

	return &types.StoredGame{
		Index:    "1",
		Board:    string(serializedBoard),
		NextTurn: string(g.NextTurn()),
		XPlayer:  alice,
		YPlayer:  bob,
	}
}

func TestCanGetAddressX(t *testing.T) {
	aliceAddr, err1 := sdk.AccAddressFromBech32(alice)
	xAddr, err2 := GetStoredGame1().XPlayerAddress()
	require.Equal(t, aliceAddr, xAddr)
	require.Nil(t, err2)
	require.Nil(t, err1)
}
