package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountUpdatePasswordSettings
func (c *TwoFaCore) AccountUpdatePasswordSettings(in *mtproto.TLAccountUpdatePasswordSettings) (*mtproto.Bool, error) {
	c.Logger.Errorf("AccountUpdatePasswordSettings - not impl")

	return mtproto.BoolTrue, nil
}
