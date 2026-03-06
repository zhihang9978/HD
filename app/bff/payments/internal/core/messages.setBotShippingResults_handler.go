package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesSetBotShippingResults
func (c *PaymentsCore) MessagesSetBotShippingResults(in *mtproto.TLMessagesSetBotShippingResults) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesSetBotShippingResults - not impl")

	return mtproto.BoolTrue, nil
}
