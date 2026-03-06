package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesReportSpam
func (c *ReportsCore) MessagesReportSpam(in *mtproto.TLMessagesReportSpam) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesReportSpam - not impl")

	return mtproto.BoolTrue, nil
}
