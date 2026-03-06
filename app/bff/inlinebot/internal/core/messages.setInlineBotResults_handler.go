package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesSetInlineBotResults
func (c *InlineBotCore) MessagesSetInlineBotResults(in *mtproto.TLMessagesSetInlineBotResults) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesSetInlineBotResults - not impl")

	return mtproto.BoolTrue, nil
}
