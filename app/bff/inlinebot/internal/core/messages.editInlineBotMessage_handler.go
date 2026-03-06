package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesEditInlineBotMessage
func (c *InlineBotCore) MessagesEditInlineBotMessage(in *mtproto.TLMessagesEditInlineBotMessage) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesEditInlineBotMessage - not impl")

	return mtproto.BoolTrue, nil
}
