package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// MessagesStartBot
func (c *DeepLinksCore) MessagesStartBot(in *mtproto.TLMessagesStartBot) (*mtproto.Updates, error) {
	c.Logger.Errorf("MessagesStartBot - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
