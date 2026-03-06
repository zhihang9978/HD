package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountSaveTheme
func (c *ThemesCore) AccountSaveTheme(in *mtproto.TLAccountSaveTheme) (*mtproto.Bool, error) {
	c.Logger.Errorf("AccountSaveTheme - not impl")

	return mtproto.BoolTrue, nil
}
