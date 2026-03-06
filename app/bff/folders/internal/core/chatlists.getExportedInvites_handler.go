package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChatlistsGetExportedInvites
func (c *FoldersCore) ChatlistsGetExportedInvites(in *mtproto.TLChatlistsGetExportedInvites) (*mtproto.Chatlists_ExportedInvites, error) {
	c.Logger.Errorf("ChatlistsGetExportedInvites - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
