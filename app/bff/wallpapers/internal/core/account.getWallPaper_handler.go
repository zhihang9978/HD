package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountGetWallPaper
func (c *WallpapersCore) AccountGetWallPaper(in *mtproto.TLAccountGetWallPaper) (*mtproto.WallPaper, error) {
	c.Logger.Errorf("AccountGetWallPaper - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
