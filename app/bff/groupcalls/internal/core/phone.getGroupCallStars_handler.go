package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneGetGroupCallStars
func (c *GroupCallsCore) PhoneGetGroupCallStars(in *mtproto.TLPhoneGetGroupCallStars) (*mtproto.Phone_GroupCallStars, error) {
	c.Logger.Errorf("PhoneGetGroupCallStars - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
