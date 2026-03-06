package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsGetInactiveChannels
func (c *ChannelsCore) ChannelsGetInactiveChannels(in *mtproto.TLChannelsGetInactiveChannels) (*mtproto.Messages_InactiveChats, error) {
	c.Logger.Errorf("ChannelsGetInactiveChannels - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
