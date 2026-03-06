package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetAvailableReactions
func (c *ReactionsCore) MessagesGetAvailableReactions(in *mtproto.TLMessagesGetAvailableReactions) (*mtproto.Messages_AvailableReactions, error) {
	c.Logger.Errorf("MessagesGetAvailableReactions - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
