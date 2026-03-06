package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChatlistsEditExportedInvite
func (c *FoldersCore) ChatlistsEditExportedInvite(in *mtproto.TLChatlistsEditExportedInvite) (*mtproto.ExportedChatlistInvite, error) {
	c.Logger.Errorf("ChatlistsEditExportedInvite - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
