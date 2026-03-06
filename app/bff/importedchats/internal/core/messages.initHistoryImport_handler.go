package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesInitHistoryImport
func (c *ImportedChatsCore) MessagesInitHistoryImport(in *mtproto.TLMessagesInitHistoryImport) (*mtproto.Messages_HistoryImport, error) {
	c.Logger.Errorf("MessagesInitHistoryImport - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
