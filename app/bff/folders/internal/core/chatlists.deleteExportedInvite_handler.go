package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChatlistsDeleteExportedInvite
func (c *FoldersCore) ChatlistsDeleteExportedInvite(in *mtproto.TLChatlistsDeleteExportedInvite) (*mtproto.Bool, error) {
	c.Logger.Errorf("ChatlistsDeleteExportedInvite - not impl")

	return mtproto.BoolTrue, nil
}
