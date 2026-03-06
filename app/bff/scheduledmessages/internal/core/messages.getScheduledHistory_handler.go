package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetScheduledHistory
func (c *ScheduledMessagesCore) MessagesGetScheduledHistory(in *mtproto.TLMessagesGetScheduledHistory) (*mtproto.Messages_Messages, error) {
	c.Logger.Errorf("MessagesGetScheduledHistory - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
