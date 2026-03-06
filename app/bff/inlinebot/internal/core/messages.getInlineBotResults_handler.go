package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetInlineBotResults
func (c *InlineBotCore) MessagesGetInlineBotResults(in *mtproto.TLMessagesGetInlineBotResults) (*mtproto.Messages_BotResults, error) {
	c.Logger.Errorf("MessagesGetInlineBotResults - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
