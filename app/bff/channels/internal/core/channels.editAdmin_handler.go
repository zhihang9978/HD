package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// ChannelsEditAdmin
func (c *ChannelsCore) ChannelsEditAdmin(in *mtproto.TLChannelsEditAdmin) (*mtproto.Updates, error) {
	c.Logger.Errorf("ChannelsEditAdmin - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
