package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsReportSpam
func (c *ReportsCore) ChannelsReportSpam(in *mtproto.TLChannelsReportSpam) (*mtproto.Bool, error) {
	c.Logger.Errorf("ChannelsReportSpam - not impl")

	return mtproto.BoolTrue, nil
}
