package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsDeleteParticipantHistory
func (c *ChannelsCore) ChannelsDeleteParticipantHistory(in *mtproto.TLChannelsDeleteParticipantHistory) (*mtproto.Messages_AffectedHistory, error) {
	c.Logger.Errorf("ChannelsDeleteParticipantHistory - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
