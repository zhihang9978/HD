package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetTopReactions
func (c *ReactionsCore) MessagesGetTopReactions(in *mtproto.TLMessagesGetTopReactions) (*mtproto.Messages_Reactions, error) {
	c.Logger.Errorf("MessagesGetTopReactions - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
