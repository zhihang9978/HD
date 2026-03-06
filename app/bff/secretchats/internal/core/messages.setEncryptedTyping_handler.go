package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesSetEncryptedTyping
func (c *SecretChatsCore) MessagesSetEncryptedTyping(in *mtproto.TLMessagesSetEncryptedTyping) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesSetEncryptedTyping - not impl")

	return mtproto.BoolTrue, nil
}
