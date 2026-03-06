package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesSetBotPrecheckoutResults
func (c *PaymentsCore) MessagesSetBotPrecheckoutResults(in *mtproto.TLMessagesSetBotPrecheckoutResults) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesSetBotPrecheckoutResults - not impl")

	return mtproto.BoolTrue, nil
}
