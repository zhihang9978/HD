package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StatsGetMessageStats
func (c *StatisticsCore) StatsGetMessageStats(in *mtproto.TLStatsGetMessageStats) (*mtproto.Stats_MessageStats, error) {
	c.Logger.Errorf("StatsGetMessageStats - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
