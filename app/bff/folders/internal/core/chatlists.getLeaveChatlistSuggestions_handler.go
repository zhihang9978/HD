package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChatlistsGetLeaveChatlistSuggestions
func (c *FoldersCore) ChatlistsGetLeaveChatlistSuggestions(in *mtproto.TLChatlistsGetLeaveChatlistSuggestions) (*mtproto.Vector_Peer, error) {
	c.Logger.Errorf("ChatlistsGetLeaveChatlistSuggestions - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
