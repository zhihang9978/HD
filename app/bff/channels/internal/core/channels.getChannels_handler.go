package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsGetChannels
func (c *ChannelsCore) ChannelsGetChannels(in *mtproto.TLChannelsGetChannels) (*mtproto.Messages_Chats, error) {
	c.Logger.Errorf("ChannelsGetChannels - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
