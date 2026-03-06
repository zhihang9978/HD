package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountGetTheme
func (c *ThemesCore) AccountGetTheme(in *mtproto.TLAccountGetTheme) (*mtproto.Theme, error) {
	c.Logger.Errorf("AccountGetTheme - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
