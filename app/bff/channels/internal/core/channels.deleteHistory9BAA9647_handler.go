package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// ChannelsDeleteHistory9BAA9647
func (c *ChannelsCore) ChannelsDeleteHistory9BAA9647(in *mtproto.TLChannelsDeleteHistory9BAA9647) (*mtproto.Updates, error) {
	c.Logger.Errorf("ChannelsDeleteHistory9BAA9647 - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
