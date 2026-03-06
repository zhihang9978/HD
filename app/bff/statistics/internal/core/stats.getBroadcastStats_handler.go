package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StatsGetBroadcastStats
func (c *StatisticsCore) StatsGetBroadcastStats(in *mtproto.TLStatsGetBroadcastStats) (*mtproto.Stats_BroadcastStats, error) {
	c.Logger.Errorf("StatsGetBroadcastStats - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
