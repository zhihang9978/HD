package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// ChannelsEditPhoto
func (c *ChannelsCore) ChannelsEditPhoto(in *mtproto.TLChannelsEditPhoto) (*mtproto.Updates, error) {
	c.Logger.Errorf("ChannelsEditPhoto - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
