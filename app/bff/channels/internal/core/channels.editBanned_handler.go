package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// ChannelsEditBanned
func (c *ChannelsCore) ChannelsEditBanned(in *mtproto.TLChannelsEditBanned) (*mtproto.Updates, error) {
	c.Logger.Errorf("ChannelsEditBanned - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
