package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesCheckHistoryImport
func (c *ImportedChatsCore) MessagesCheckHistoryImport(in *mtproto.TLMessagesCheckHistoryImport) (*mtproto.Messages_HistoryImportParsed, error) {
	c.Logger.Errorf("MessagesCheckHistoryImport - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
