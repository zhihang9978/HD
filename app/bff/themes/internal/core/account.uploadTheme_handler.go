package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountUploadTheme
func (c *ThemesCore) AccountUploadTheme(in *mtproto.TLAccountUploadTheme) (*mtproto.Document, error) {
	c.Logger.Errorf("AccountUploadTheme - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
