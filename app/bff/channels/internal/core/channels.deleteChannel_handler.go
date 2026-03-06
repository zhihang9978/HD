package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// ChannelsDeleteChannel
func (c *ChannelsCore) ChannelsDeleteChannel(in *mtproto.TLChannelsDeleteChannel) (*mtproto.Updates, error) {
	c.Logger.Errorf("ChannelsDeleteChannel - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
