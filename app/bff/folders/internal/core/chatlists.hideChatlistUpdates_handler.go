package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChatlistsHideChatlistUpdates
func (c *FoldersCore) ChatlistsHideChatlistUpdates(in *mtproto.TLChatlistsHideChatlistUpdates) (*mtproto.Bool, error) {
	c.Logger.Errorf("ChatlistsHideChatlistUpdates - not impl")

	return mtproto.BoolTrue, nil
}
