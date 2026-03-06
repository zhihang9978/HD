package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountReportPeer
func (c *ReportsCore) AccountReportPeer(in *mtproto.TLAccountReportPeer) (*mtproto.Bool, error) {
	c.Logger.Errorf("AccountReportPeer - not impl")

	return mtproto.BoolTrue, nil
}
