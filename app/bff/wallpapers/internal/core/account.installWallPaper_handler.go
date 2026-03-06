package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountInstallWallPaper
func (c *WallpapersCore) AccountInstallWallPaper(in *mtproto.TLAccountInstallWallPaper) (*mtproto.Bool, error) {
	c.Logger.Errorf("AccountInstallWallPaper - not impl")

	return mtproto.BoolTrue, nil
}
