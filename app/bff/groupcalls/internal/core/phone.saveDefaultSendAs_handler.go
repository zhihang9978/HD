package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneSaveDefaultSendAs
func (c *GroupCallsCore) PhoneSaveDefaultSendAs(in *mtproto.TLPhoneSaveDefaultSendAs) (*mtproto.Bool, error) {
	c.Logger.Errorf("PhoneSaveDefaultSendAs - not impl")

	return mtproto.BoolTrue, nil
}
