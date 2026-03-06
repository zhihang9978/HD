package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneSaveDefaultGroupCallJoinAs
func (c *GroupCallsCore) PhoneSaveDefaultGroupCallJoinAs(in *mtproto.TLPhoneSaveDefaultGroupCallJoinAs) (*mtproto.Bool, error) {
	c.Logger.Errorf("PhoneSaveDefaultGroupCallJoinAs - not impl")

	return mtproto.BoolTrue, nil
}
