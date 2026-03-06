package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// ChannelsToggleSignatures
func (c *ChannelsCore) ChannelsToggleSignatures(in *mtproto.TLChannelsToggleSignatures) (*mtproto.Updates, error) {
	c.Logger.Errorf("ChannelsToggleSignatures - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
