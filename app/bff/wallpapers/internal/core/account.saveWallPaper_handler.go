package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountSaveWallPaper
func (c *WallpapersCore) AccountSaveWallPaper(in *mtproto.TLAccountSaveWallPaper) (*mtproto.Bool, error) {
	c.Logger.Errorf("AccountSaveWallPaper - not impl")

	return mtproto.BoolTrue, nil
}
