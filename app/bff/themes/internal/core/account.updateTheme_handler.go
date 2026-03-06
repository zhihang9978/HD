package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountUpdateTheme
func (c *ThemesCore) AccountUpdateTheme(in *mtproto.TLAccountUpdateTheme) (*mtproto.Theme, error) {
	c.Logger.Errorf("AccountUpdateTheme - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
