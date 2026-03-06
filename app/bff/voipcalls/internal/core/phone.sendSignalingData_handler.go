package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneSendSignalingData
func (c *VoipCallsCore) PhoneSendSignalingData(in *mtproto.TLPhoneSendSignalingData) (*mtproto.Bool, error) {
	c.Logger.Errorf("PhoneSendSignalingData - not impl")

	return mtproto.BoolTrue, nil
}
