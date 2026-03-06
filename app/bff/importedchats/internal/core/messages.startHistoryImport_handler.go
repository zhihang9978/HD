package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesStartHistoryImport
func (c *ImportedChatsCore) MessagesStartHistoryImport(in *mtproto.TLMessagesStartHistoryImport) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesStartHistoryImport - not impl")

	return mtproto.BoolTrue, nil
}
