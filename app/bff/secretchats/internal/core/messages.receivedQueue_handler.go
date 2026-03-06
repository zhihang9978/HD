package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesReceivedQueue
func (c *SecretChatsCore) MessagesReceivedQueue(in *mtproto.TLMessagesReceivedQueue) (*mtproto.Vector_Long, error) {
	c.Logger.Errorf("MessagesReceivedQueue - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
