package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountConfirmPasswordEmail
func (c *TwoFaCore) AccountConfirmPasswordEmail(in *mtproto.TLAccountConfirmPasswordEmail) (*mtproto.Bool, error) {
	c.Logger.Errorf("AccountConfirmPasswordEmail - not impl")

	return mtproto.BoolTrue, nil
}
