package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesReportEncryptedSpam
func (c *ReportsCore) MessagesReportEncryptedSpam(in *mtproto.TLMessagesReportEncryptedSpam) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesReportEncryptedSpam - not impl")

	return mtproto.BoolTrue, nil
}
