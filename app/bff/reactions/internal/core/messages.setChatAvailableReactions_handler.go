package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// MessagesSetChatAvailableReactions
func (c *ReactionsCore) MessagesSetChatAvailableReactions(in *mtproto.TLMessagesSetChatAvailableReactions) (*mtproto.Updates, error) {
	c.Logger.Errorf("MessagesSetChatAvailableReactions - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
