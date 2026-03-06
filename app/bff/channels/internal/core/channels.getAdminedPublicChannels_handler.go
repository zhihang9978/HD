package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsGetAdminedPublicChannels
func (c *ChannelsCore) ChannelsGetAdminedPublicChannels(in *mtproto.TLChannelsGetAdminedPublicChannels) (*mtproto.Messages_Chats, error) {
	c.Logger.Errorf("ChannelsGetAdminedPublicChannels - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
