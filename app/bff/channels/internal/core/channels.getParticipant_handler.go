package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsGetParticipant
func (c *ChannelsCore) ChannelsGetParticipant(in *mtproto.TLChannelsGetParticipant) (*mtproto.Channels_ChannelParticipant, error) {
	c.Logger.Errorf("ChannelsGetParticipant - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
