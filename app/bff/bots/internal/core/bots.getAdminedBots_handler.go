package core

import (
	"github.com/teamgram/proto/mtproto"
)

// BotsGetAdminedBots
func (c *BotsCore) BotsGetAdminedBots(in *mtproto.TLBotsGetAdminedBots) (*mtproto.Vector_User, error) {
	c.Logger.Errorf("BotsGetAdminedBots - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
