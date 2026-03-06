package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneSendGroupCallMessage87893014
func (c *GroupCallsCore) PhoneSendGroupCallMessage87893014(in *mtproto.TLPhoneSendGroupCallMessage87893014) (*mtproto.Bool, error) {
	c.Logger.Errorf("PhoneSendGroupCallMessage87893014 - not impl")

	return mtproto.BoolTrue, nil
}
