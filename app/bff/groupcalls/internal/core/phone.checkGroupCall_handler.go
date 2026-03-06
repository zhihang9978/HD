package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneCheckGroupCall
func (c *GroupCallsCore) PhoneCheckGroupCall(in *mtproto.TLPhoneCheckGroupCall) (*mtproto.Vector_Int, error) {
	c.Logger.Errorf("PhoneCheckGroupCall - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
