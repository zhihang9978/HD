package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StatsGetMessagePublicForwards5630281B
func (c *StatisticsCore) StatsGetMessagePublicForwards5630281B(in *mtproto.TLStatsGetMessagePublicForwards5630281B) (*mtproto.Messages_Messages, error) {
	c.Logger.Errorf("StatsGetMessagePublicForwards5630281B - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
