package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/palomachain/paloma/v2/util/libmeta"
	vtypes "github.com/palomachain/paloma/v2/x/valset/types"
)

const TypeMsgExecuteJob = "execute_job"

var _ sdk.Msg = &MsgExecuteJob{}

func NewMsgExecuteJob(creator string, jobID string, payload []byte) *MsgExecuteJob {
	return &MsgExecuteJob{
		JobID:   jobID,
		Payload: payload,
		Metadata: vtypes.MsgMetadata{
			Creator: creator,
		},
	}
}

func (msg *MsgExecuteJob) Route() string {
	return RouterKey
}

func (msg *MsgExecuteJob) Type() string {
	return TypeMsgExecuteJob
}

func (msg *MsgExecuteJob) GetSigners() []sdk.AccAddress {
	return libmeta.GetSigners(msg)
}

func (msg *MsgExecuteJob) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgExecuteJob) ValidateBasic() error {
	return libmeta.ValidateBasic(msg)
}
