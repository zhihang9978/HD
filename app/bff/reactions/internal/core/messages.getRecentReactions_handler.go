package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetRecentReactions
func (c *ReactionsCore) MessagesGetRecentReactions(in *mtproto.TLMessagesGetRecentReactions) (*mtproto.Messages_Reactions, error) {
	c.Logger.Errorf("MessagesGetRecentReactions - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
