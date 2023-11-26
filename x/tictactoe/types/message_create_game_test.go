package types

import (
	"github.com/alice/checkers/x/tictactoe/testutil"
	"testing"

	"github.com/alice/checkers/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateGame_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateGame
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateGame{
				Creator: "invalid_address",
				XPlayer: testutil.Alice,
				OPlayer: testutil.Bob,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateGame{
				Creator: sample.AccAddress(),
				XPlayer: testutil.Alice,
				OPlayer: testutil.Bob,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
