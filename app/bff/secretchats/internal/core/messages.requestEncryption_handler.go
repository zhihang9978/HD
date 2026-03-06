package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesRequestEncryption
func (c *SecretChatsCore) MessagesRequestEncryption(in *mtproto.TLMessagesRequestEncryption) (*mtproto.EncryptedChat, error) {
	c.Logger.Errorf("MessagesRequestEncryption - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
