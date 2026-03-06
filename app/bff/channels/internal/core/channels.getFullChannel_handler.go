package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsGetFullChannel
func (c *ChannelsCore) ChannelsGetFullChannel(in *mtproto.TLChannelsGetFullChannel) (*mtproto.Messages_ChatFull, error) {
	c.Logger.Errorf("ChannelsGetFullChannel - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
