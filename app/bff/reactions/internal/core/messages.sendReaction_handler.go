package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// MessagesSendReaction
func (c *ReactionsCore) MessagesSendReaction(in *mtproto.TLMessagesSendReaction) (*mtproto.Updates, error) {
	c.Logger.Errorf("MessagesSendReaction - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
