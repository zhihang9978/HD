package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesReadDiscussion
func (c *MessageThreadsCore) MessagesReadDiscussion(in *mtproto.TLMessagesReadDiscussion) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesReadDiscussion - not impl")

	return mtproto.BoolTrue, nil
}
