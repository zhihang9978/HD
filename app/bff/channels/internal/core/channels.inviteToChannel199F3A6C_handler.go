package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// ChannelsInviteToChannel199F3A6C
func (c *ChannelsCore) ChannelsInviteToChannel199F3A6C(in *mtproto.TLChannelsInviteToChannel199F3A6C) (*mtproto.Updates, error) {
	c.Logger.Errorf("ChannelsInviteToChannel199F3A6C - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
