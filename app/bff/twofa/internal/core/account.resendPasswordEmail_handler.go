package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountResendPasswordEmail
func (c *TwoFaCore) AccountResendPasswordEmail(in *mtproto.TLAccountResendPasswordEmail) (*mtproto.Bool, error) {
	c.Logger.Errorf("AccountResendPasswordEmail - not impl")

	return mtproto.BoolTrue, nil
}
