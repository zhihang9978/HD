package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountGetThemes
func (c *ThemesCore) AccountGetThemes(in *mtproto.TLAccountGetThemes) (*mtproto.Account_Themes, error) {
	c.Logger.Errorf("AccountGetThemes - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
