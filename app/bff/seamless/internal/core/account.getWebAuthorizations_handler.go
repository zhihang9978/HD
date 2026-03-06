package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountGetWebAuthorizations
func (c *SeamlessCore) AccountGetWebAuthorizations(in *mtproto.TLAccountGetWebAuthorizations) (*mtproto.Account_WebAuthorizations, error) {
	c.Logger.Errorf("AccountGetWebAuthorizations - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
