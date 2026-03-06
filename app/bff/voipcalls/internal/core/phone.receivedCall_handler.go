package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneReceivedCall
func (c *VoipCallsCore) PhoneReceivedCall(in *mtproto.TLPhoneReceivedCall) (*mtproto.Bool, error) {
	c.Logger.Errorf("PhoneReceivedCall - not impl")

	return mtproto.BoolTrue, nil
}
