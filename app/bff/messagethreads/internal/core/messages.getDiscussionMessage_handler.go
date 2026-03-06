package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetDiscussionMessage
func (c *MessageThreadsCore) MessagesGetDiscussionMessage(in *mtproto.TLMessagesGetDiscussionMessage) (*mtproto.Messages_DiscussionMessage, error) {
	c.Logger.Errorf("MessagesGetDiscussionMessage - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
