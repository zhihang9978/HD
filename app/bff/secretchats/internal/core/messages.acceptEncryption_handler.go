package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesAcceptEncryption
func (c *SecretChatsCore) MessagesAcceptEncryption(in *mtproto.TLMessagesAcceptEncryption) (*mtproto.EncryptedChat, error) {
	c.Logger.Errorf("MessagesAcceptEncryption - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
