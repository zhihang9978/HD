package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChatlistsExportChatlistInvite
func (c *FoldersCore) ChatlistsExportChatlistInvite(in *mtproto.TLChatlistsExportChatlistInvite) (*mtproto.Chatlists_ExportedChatlistInvite, error) {
	c.Logger.Errorf("ChatlistsExportChatlistInvite - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
