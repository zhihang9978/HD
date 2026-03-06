package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetBotCallbackAnswer
func (c *InlineBotCore) MessagesGetBotCallbackAnswer(in *mtproto.TLMessagesGetBotCallbackAnswer) (*mtproto.Messages_BotCallbackAnswer, error) {
	c.Logger.Errorf("MessagesGetBotCallbackAnswer - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
