package core

import (
	"github.com/teamgram/proto/mtproto"
)

// BotsResetBotCommands
func (c *BotsCore) BotsResetBotCommands(in *mtproto.TLBotsResetBotCommands) (*mtproto.Bool, error) {
	c.Logger.Errorf("BotsResetBotCommands - not impl")

	return mtproto.BoolTrue, nil
}
