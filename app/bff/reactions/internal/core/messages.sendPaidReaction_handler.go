package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// MessagesSendPaidReaction
func (c *ReactionsCore) MessagesSendPaidReaction(in *mtproto.TLMessagesSendPaidReaction) (*mtproto.Updates, error) {
	c.Logger.Errorf("MessagesSendPaidReaction - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
