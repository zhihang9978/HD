package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountGetPasswordSettings
func (c *TwoFaCore) AccountGetPasswordSettings(in *mtproto.TLAccountGetPasswordSettings) (*mtproto.Account_PasswordSettings, error) {
	c.Logger.Errorf("AccountGetPasswordSettings - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
