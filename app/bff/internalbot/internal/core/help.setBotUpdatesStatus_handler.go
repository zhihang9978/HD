package core

import (
	"github.com/teamgram/proto/mtproto"
)

// HelpSetBotUpdatesStatus
func (c *InternalBotCore) HelpSetBotUpdatesStatus(in *mtproto.TLHelpSetBotUpdatesStatus) (*mtproto.Bool, error) {
	c.Logger.Errorf("HelpSetBotUpdatesStatus - not impl")

	return mtproto.BoolTrue, nil
}
