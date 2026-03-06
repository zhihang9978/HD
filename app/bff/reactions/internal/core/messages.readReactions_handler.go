package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesReadReactions
func (c *ReactionsCore) MessagesReadReactions(in *mtproto.TLMessagesReadReactions) (*mtproto.Messages_AffectedHistory, error) {
	c.Logger.Errorf("MessagesReadReactions - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
