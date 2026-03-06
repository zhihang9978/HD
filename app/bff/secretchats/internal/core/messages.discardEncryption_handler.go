package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesDiscardEncryption
func (c *SecretChatsCore) MessagesDiscardEncryption(in *mtproto.TLMessagesDiscardEncryption) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesDiscardEncryption - not impl")

	return mtproto.BoolTrue, nil
}
