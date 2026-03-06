package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// ChatlistsJoinChatlistUpdates
func (c *FoldersCore) ChatlistsJoinChatlistUpdates(in *mtproto.TLChatlistsJoinChatlistUpdates) (*mtproto.Updates, error) {
	c.Logger.Errorf("ChatlistsJoinChatlistUpdates - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
