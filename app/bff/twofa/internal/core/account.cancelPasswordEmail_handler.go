package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountCancelPasswordEmail
func (c *TwoFaCore) AccountCancelPasswordEmail(in *mtproto.TLAccountCancelPasswordEmail) (*mtproto.Bool, error) {
	c.Logger.Errorf("AccountCancelPasswordEmail - not impl")

	return mtproto.BoolTrue, nil
}
