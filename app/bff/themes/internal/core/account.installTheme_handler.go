package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountInstallTheme
func (c *ThemesCore) AccountInstallTheme(in *mtproto.TLAccountInstallTheme) (*mtproto.Bool, error) {
	c.Logger.Errorf("AccountInstallTheme - not impl")

	return mtproto.BoolTrue, nil
}
