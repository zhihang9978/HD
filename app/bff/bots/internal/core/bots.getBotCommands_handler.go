package core

import (
	"github.com/teamgram/proto/mtproto"
)

// BotsGetBotCommands
func (c *BotsCore) BotsGetBotCommands(in *mtproto.TLBotsGetBotCommands) (*mtproto.Vector_BotCommand, error) {
	c.Logger.Errorf("BotsGetBotCommands - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
