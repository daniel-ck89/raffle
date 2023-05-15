package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"raffle/testutil/sample"
)

func TestMsgCreateSimpleRaffle_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateSimpleRaffle
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateSimpleRaffle{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateSimpleRaffle{
				Creator: sample.AccAddress(),
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
