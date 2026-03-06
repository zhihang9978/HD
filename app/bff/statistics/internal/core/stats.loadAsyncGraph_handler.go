package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StatsLoadAsyncGraph
func (c *StatisticsCore) StatsLoadAsyncGraph(in *mtproto.TLStatsLoadAsyncGraph) (*mtproto.StatsGraph, error) {
	c.Logger.Errorf("StatsLoadAsyncGraph - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
