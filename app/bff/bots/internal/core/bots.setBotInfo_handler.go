package core

import (
	"github.com/teamgram/proto/mtproto"
)

// BotsSetBotInfo
func (c *BotsCore) BotsSetBotInfo(in *mtproto.TLBotsSetBotInfo) (*mtproto.Bool, error) {
	c.Logger.Errorf("BotsSetBotInfo - not impl")

	return mtproto.BoolTrue, nil
}
