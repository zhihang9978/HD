package core

import (
	"github.com/teamgram/proto/mtproto"
)

// BotsGetBotInfoDCD914FD
func (c *BotsCore) BotsGetBotInfoDCD914FD(in *mtproto.TLBotsGetBotInfoDCD914FD) (*mtproto.Bots_BotInfo, error) {
	c.Logger.Errorf("BotsGetBotInfoDCD914FD - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
