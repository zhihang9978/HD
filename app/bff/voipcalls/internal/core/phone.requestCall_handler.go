package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneRequestCall
func (c *VoipCallsCore) PhoneRequestCall(in *mtproto.TLPhoneRequestCall) (*mtproto.Phone_PhoneCall, error) {
	c.Logger.Errorf("PhoneRequestCall - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
