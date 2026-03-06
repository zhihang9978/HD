package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetScheduledMessages
func (c *ScheduledMessagesCore) MessagesGetScheduledMessages(in *mtproto.TLMessagesGetScheduledMessages) (*mtproto.Messages_Messages, error) {
	c.Logger.Errorf("MessagesGetScheduledMessages - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
