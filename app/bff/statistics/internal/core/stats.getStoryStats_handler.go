package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StatsGetStoryStats
func (c *StatisticsCore) StatsGetStoryStats(in *mtproto.TLStatsGetStoryStats) (*mtproto.Stats_StoryStats, error) {
	c.Logger.Errorf("StatsGetStoryStats - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
