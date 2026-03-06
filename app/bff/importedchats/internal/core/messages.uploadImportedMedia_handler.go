package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesUploadImportedMedia
func (c *ImportedChatsCore) MessagesUploadImportedMedia(in *mtproto.TLMessagesUploadImportedMedia) (*mtproto.MessageMedia, error) {
	c.Logger.Errorf("MessagesUploadImportedMedia - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
