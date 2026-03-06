package core

import (
	"github.com/teamgram/proto/mtproto"
)

// BotsSendCustomRequest
func (c *InternalBotCore) BotsSendCustomRequest(in *mtproto.TLBotsSendCustomRequest) (*mtproto.DataJSON, error) {
	c.Logger.Errorf("BotsSendCustomRequest - not impl")

	return mtproto.MakeTLDataJSON(&mtproto.DataJSON{Data: "{}"}).To_DataJSON(), nil
}
