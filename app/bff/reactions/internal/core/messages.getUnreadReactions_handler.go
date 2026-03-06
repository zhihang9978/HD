package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetUnreadReactions
func (c *ReactionsCore) MessagesGetUnreadReactions(in *mtproto.TLMessagesGetUnreadReactions) (*mtproto.Messages_Messages, error) {
	c.Logger.Errorf("MessagesGetUnreadReactions - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
