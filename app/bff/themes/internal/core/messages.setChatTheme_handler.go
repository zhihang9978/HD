package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// MessagesSetChatTheme
func (c *ThemesCore) MessagesSetChatTheme(in *mtproto.TLMessagesSetChatTheme) (*mtproto.Updates, error) {
	c.Logger.Errorf("MessagesSetChatTheme - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
