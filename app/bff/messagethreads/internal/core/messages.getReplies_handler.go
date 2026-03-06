package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetReplies
func (c *MessageThreadsCore) MessagesGetReplies(in *mtproto.TLMessagesGetReplies) (*mtproto.Messages_Messages, error) {
	c.Logger.Errorf("MessagesGetReplies - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
