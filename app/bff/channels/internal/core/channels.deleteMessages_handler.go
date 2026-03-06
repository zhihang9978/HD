package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsDeleteMessages
func (c *ChannelsCore) ChannelsDeleteMessages(in *mtproto.TLChannelsDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	c.Logger.Errorf("ChannelsDeleteMessages - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
