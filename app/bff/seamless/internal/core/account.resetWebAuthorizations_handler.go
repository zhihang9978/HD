package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountResetWebAuthorizations
func (c *SeamlessCore) AccountResetWebAuthorizations(in *mtproto.TLAccountResetWebAuthorizations) (*mtproto.Bool, error) {
	c.Logger.Errorf("AccountResetWebAuthorizations - not impl")

	return mtproto.BoolTrue, nil
}
