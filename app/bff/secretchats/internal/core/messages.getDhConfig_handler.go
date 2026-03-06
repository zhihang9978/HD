package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetDhConfig
func (c *SecretChatsCore) MessagesGetDhConfig(in *mtproto.TLMessagesGetDhConfig) (*mtproto.Messages_DhConfig, error) {
	c.Logger.Errorf("MessagesGetDhConfig - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
