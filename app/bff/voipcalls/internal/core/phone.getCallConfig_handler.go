package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneGetCallConfig
func (c *VoipCallsCore) PhoneGetCallConfig(in *mtproto.TLPhoneGetCallConfig) (*mtproto.DataJSON, error) {
	c.Logger.Errorf("PhoneGetCallConfig - not impl")

	return mtproto.MakeTLDataJSON(&mtproto.DataJSON{Data: "{}"}).To_DataJSON(), nil
}
