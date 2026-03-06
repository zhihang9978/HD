package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// ChannelsCreateChannel
func (c *ChannelsCore) ChannelsCreateChannel(in *mtproto.TLChannelsCreateChannel) (*mtproto.Updates, error) {
	c.Logger.Errorf("ChannelsCreateChannel - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
