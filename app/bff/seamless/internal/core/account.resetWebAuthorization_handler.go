package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountResetWebAuthorization
func (c *SeamlessCore) AccountResetWebAuthorization(in *mtproto.TLAccountResetWebAuthorization) (*mtproto.Bool, error) {
	c.Logger.Errorf("AccountResetWebAuthorization - not impl")

	return mtproto.BoolTrue, nil
}
