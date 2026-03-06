package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetMessageReactionsList
func (c *ReactionsCore) MessagesGetMessageReactionsList(in *mtproto.TLMessagesGetMessageReactionsList) (*mtproto.Messages_MessageReactionsList, error) {
	c.Logger.Errorf("MessagesGetMessageReactionsList - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
