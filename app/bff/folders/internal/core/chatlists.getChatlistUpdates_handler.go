package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChatlistsGetChatlistUpdates
func (c *FoldersCore) ChatlistsGetChatlistUpdates(in *mtproto.TLChatlistsGetChatlistUpdates) (*mtproto.Chatlists_ChatlistUpdates, error) {
	c.Logger.Errorf("ChatlistsGetChatlistUpdates - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
