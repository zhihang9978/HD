package core

import (
	"github.com/teamgram/proto/mtproto"
)

// BotsAnswerWebhookJSONQuery
func (c *InternalBotCore) BotsAnswerWebhookJSONQuery(in *mtproto.TLBotsAnswerWebhookJSONQuery) (*mtproto.Bool, error) {
	c.Logger.Errorf("BotsAnswerWebhookJSONQuery - not impl")

	return mtproto.BoolTrue, nil
}
