package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesSetBotCallbackAnswer
func (c *InlineBotCore) MessagesSetBotCallbackAnswer(in *mtproto.TLMessagesSetBotCallbackAnswer) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesSetBotCallbackAnswer - not impl")

	return mtproto.BoolTrue, nil
}
