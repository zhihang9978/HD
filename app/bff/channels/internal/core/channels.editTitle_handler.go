package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// ChannelsEditTitle
func (c *ChannelsCore) ChannelsEditTitle(in *mtproto.TLChannelsEditTitle) (*mtproto.Updates, error) {
	c.Logger.Errorf("ChannelsEditTitle - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
