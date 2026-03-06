package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// ChannelsLeaveChannel
func (c *ChannelsCore) ChannelsLeaveChannel(in *mtproto.TLChannelsLeaveChannel) (*mtproto.Updates, error) {
	c.Logger.Errorf("ChannelsLeaveChannel - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
