package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneSaveCallLog
func (c *VoipCallsCore) PhoneSaveCallLog(in *mtproto.TLPhoneSaveCallLog) (*mtproto.Bool, error) {
	c.Logger.Errorf("PhoneSaveCallLog - not impl")

	return mtproto.BoolTrue, nil
}
