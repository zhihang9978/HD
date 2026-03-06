package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneSaveCallDebug
func (c *VoipCallsCore) PhoneSaveCallDebug(in *mtproto.TLPhoneSaveCallDebug) (*mtproto.Bool, error) {
	c.Logger.Errorf("PhoneSaveCallDebug - not impl")

	return mtproto.BoolTrue, nil
}
