package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountGetPassword
func (c *TwoFaCore) AccountGetPassword(in *mtproto.TLAccountGetPassword) (*mtproto.Account_Password, error) {
	c.Logger.Errorf("AccountGetPassword - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
