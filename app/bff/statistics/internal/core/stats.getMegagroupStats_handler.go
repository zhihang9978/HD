package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StatsGetMegagroupStats
func (c *StatisticsCore) StatsGetMegagroupStats(in *mtproto.TLStatsGetMegagroupStats) (*mtproto.Stats_MegagroupStats, error) {
	c.Logger.Errorf("StatsGetMegagroupStats - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
