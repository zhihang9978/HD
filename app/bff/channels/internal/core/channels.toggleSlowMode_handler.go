package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// ChannelsToggleSlowMode
func (c *ChannelsCore) ChannelsToggleSlowMode(in *mtproto.TLChannelsToggleSlowMode) (*mtproto.Updates, error) {
	c.Logger.Errorf("ChannelsToggleSlowMode - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
