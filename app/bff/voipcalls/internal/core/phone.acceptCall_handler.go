package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneAcceptCall
func (c *VoipCallsCore) PhoneAcceptCall(in *mtproto.TLPhoneAcceptCall) (*mtproto.Phone_PhoneCall, error) {
	c.Logger.Errorf("PhoneAcceptCall - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
