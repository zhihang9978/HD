package core

import (
	"github.com/teamgram/proto/mtproto"
)

// BotsSetBotCommands
func (c *BotsCore) BotsSetBotCommands(in *mtproto.TLBotsSetBotCommands) (*mtproto.Bool, error) {
	c.Logger.Errorf("BotsSetBotCommands - not impl")

	return mtproto.BoolTrue, nil
}
