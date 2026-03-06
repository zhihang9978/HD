package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetPaidReactionPrivacy
func (c *ReactionsCore) MessagesGetPaidReactionPrivacy(in *mtproto.TLMessagesGetPaidReactionPrivacy) (*mtproto.Updates, error) {
	c.Logger.Errorf("MessagesGetPaidReactionPrivacy - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
