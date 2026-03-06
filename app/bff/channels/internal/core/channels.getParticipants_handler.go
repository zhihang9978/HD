package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsGetParticipants
func (c *ChannelsCore) ChannelsGetParticipants(in *mtproto.TLChannelsGetParticipants) (*mtproto.Channels_ChannelParticipants, error) {
	c.Logger.Errorf("ChannelsGetParticipants - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
