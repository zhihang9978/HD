package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountDeclinePasswordReset
func (c *TwoFaCore) AccountDeclinePasswordReset(in *mtproto.TLAccountDeclinePasswordReset) (*mtproto.Bool, error) {
	c.Logger.Errorf("AccountDeclinePasswordReset - not impl")

	return mtproto.BoolTrue, nil
}
