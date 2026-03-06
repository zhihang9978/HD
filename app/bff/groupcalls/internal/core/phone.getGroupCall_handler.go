package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneGetGroupCall
func (c *GroupCallsCore) PhoneGetGroupCall(in *mtproto.TLPhoneGetGroupCall) (*mtproto.Phone_GroupCall, error) {
	c.Logger.Errorf("PhoneGetGroupCall - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
