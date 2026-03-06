package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesReportReaction
func (c *ReactionsCore) MessagesReportReaction(in *mtproto.TLMessagesReportReaction) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesReportReaction - not impl")

	return mtproto.BoolTrue, nil
}
