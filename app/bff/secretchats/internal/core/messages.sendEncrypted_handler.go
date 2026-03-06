package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesSendEncrypted
func (c *SecretChatsCore) MessagesSendEncrypted(in *mtproto.TLMessagesSendEncrypted) (*mtproto.Messages_SentEncryptedMessage, error) {
	c.Logger.Errorf("MessagesSendEncrypted - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
