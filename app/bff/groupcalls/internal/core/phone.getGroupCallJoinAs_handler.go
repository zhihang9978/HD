package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneGetGroupCallJoinAs
func (c *GroupCallsCore) PhoneGetGroupCallJoinAs(in *mtproto.TLPhoneGetGroupCallJoinAs) (*mtproto.Phone_JoinAsPeers, error) {
	c.Logger.Errorf("PhoneGetGroupCallJoinAs - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
