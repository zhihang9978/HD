package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesSendEncryptedFile
func (c *SecretChatsCore) MessagesSendEncryptedFile(in *mtproto.TLMessagesSendEncryptedFile) (*mtproto.Messages_SentEncryptedMessage, error) {
	c.Logger.Errorf("MessagesSendEncryptedFile - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
