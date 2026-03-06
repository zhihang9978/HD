package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountCreateTheme
func (c *ThemesCore) AccountCreateTheme(in *mtproto.TLAccountCreateTheme) (*mtproto.Theme, error) {
	c.Logger.Errorf("AccountCreateTheme - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
