package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StatsGetMessagePublicForwards5F150144
func (c *StatisticsCore) StatsGetMessagePublicForwards5F150144(in *mtproto.TLStatsGetMessagePublicForwards5F150144) (*mtproto.Stats_PublicForwards, error) {
	c.Logger.Errorf("StatsGetMessagePublicForwards5F150144 - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
