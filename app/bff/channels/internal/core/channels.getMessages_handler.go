package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsGetMessages
func (c *ChannelsCore) ChannelsGetMessages(in *mtproto.TLChannelsGetMessages) (*mtproto.Messages_Messages, error) {
	c.Logger.Errorf("ChannelsGetMessages - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
