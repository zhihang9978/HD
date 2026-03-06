package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneConfirmCall
func (c *VoipCallsCore) PhoneConfirmCall(in *mtproto.TLPhoneConfirmCall) (*mtproto.Phone_PhoneCall, error) {
	c.Logger.Errorf("PhoneConfirmCall - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
