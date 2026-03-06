package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// ChannelsJoinChannel
func (c *ChannelsCore) ChannelsJoinChannel(in *mtproto.TLChannelsJoinChannel) (*mtproto.Updates, error) {
	c.Logger.Errorf("ChannelsJoinChannel - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
