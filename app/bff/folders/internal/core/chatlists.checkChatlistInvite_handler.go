package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChatlistsCheckChatlistInvite
func (c *FoldersCore) ChatlistsCheckChatlistInvite(in *mtproto.TLChatlistsCheckChatlistInvite) (*mtproto.Chatlists_ChatlistInvite, error) {
	c.Logger.Errorf("ChatlistsCheckChatlistInvite - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
