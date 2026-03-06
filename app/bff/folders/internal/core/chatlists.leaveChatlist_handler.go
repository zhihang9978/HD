package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// ChatlistsLeaveChatlist
func (c *FoldersCore) ChatlistsLeaveChatlist(in *mtproto.TLChatlistsLeaveChatlist) (*mtproto.Updates, error) {
	c.Logger.Errorf("ChatlistsLeaveChatlist - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
