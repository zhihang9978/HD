package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneGetGroupCallStreamChannels
func (c *GroupCallsCore) PhoneGetGroupCallStreamChannels(in *mtproto.TLPhoneGetGroupCallStreamChannels) (*mtproto.Phone_GroupCallStreamChannels, error) {
	c.Logger.Errorf("PhoneGetGroupCallStreamChannels - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
