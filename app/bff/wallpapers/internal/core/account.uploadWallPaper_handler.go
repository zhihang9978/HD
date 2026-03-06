package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountUploadWallPaper
func (c *WallpapersCore) AccountUploadWallPaper(in *mtproto.TLAccountUploadWallPaper) (*mtproto.WallPaper, error) {
	c.Logger.Errorf("AccountUploadWallPaper - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
