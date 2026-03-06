package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// MessagesSetChatWallPaper
func (c *WallpapersCore) MessagesSetChatWallPaper(in *mtproto.TLMessagesSetChatWallPaper) (*mtproto.Updates, error) {
	c.Logger.Errorf("MessagesSetChatWallPaper - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
