package types

import (
	"testing"

	"github.com/CudoVentures/cudos-node/testutil/sample"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgPublishNft_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgPublishNft
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgPublishNft{
				TokenId: "1",
				DenomId: "1",
				Price:   sdk.NewCoin("acudos", sdk.NewInt(1000)),
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgPublishNft{
				TokenId: "1",
				DenomId: "1",
				Price:   sdk.NewCoin("acudos", sdk.NewInt(999)),
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
