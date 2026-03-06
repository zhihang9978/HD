package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesCheckHistoryImportPeer
func (c *ImportedChatsCore) MessagesCheckHistoryImportPeer(in *mtproto.TLMessagesCheckHistoryImportPeer) (*mtproto.Messages_CheckedHistoryImportPeer, error) {
	c.Logger.Errorf("MessagesCheckHistoryImportPeer - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
