package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesSendEncryptedService
func (c *SecretChatsCore) MessagesSendEncryptedService(in *mtproto.TLMessagesSendEncryptedService) (*mtproto.Messages_SentEncryptedMessage, error) {
	c.Logger.Errorf("MessagesSendEncryptedService - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
