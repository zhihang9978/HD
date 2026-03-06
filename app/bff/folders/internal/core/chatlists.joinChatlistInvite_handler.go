package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// ChatlistsJoinChatlistInvite
func (c *FoldersCore) ChatlistsJoinChatlistInvite(in *mtproto.TLChatlistsJoinChatlistInvite) (*mtproto.Updates, error) {
	c.Logger.Errorf("ChatlistsJoinChatlistInvite - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
