package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesReadEncryptedHistory
func (c *SecretChatsCore) MessagesReadEncryptedHistory(in *mtproto.TLMessagesReadEncryptedHistory) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesReadEncryptedHistory - not impl")

	return mtproto.BoolTrue, nil
}
