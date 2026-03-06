package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StatsGetStoryPublicForwards
func (c *StatisticsCore) StatsGetStoryPublicForwards(in *mtproto.TLStatsGetStoryPublicForwards) (*mtproto.Stats_PublicForwards, error) {
	c.Logger.Errorf("StatsGetStoryPublicForwards - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
