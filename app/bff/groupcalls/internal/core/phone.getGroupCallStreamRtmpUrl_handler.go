package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneGetGroupCallStreamRtmpUrl
func (c *GroupCallsCore) PhoneGetGroupCallStreamRtmpUrl(in *mtproto.TLPhoneGetGroupCallStreamRtmpUrl) (*mtproto.Phone_GroupCallStreamRtmpUrl, error) {
	c.Logger.Errorf("PhoneGetGroupCallStreamRtmpUrl - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
