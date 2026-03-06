package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountGetWallPapers
func (c *WallpapersCore) AccountGetWallPapers(in *mtproto.TLAccountGetWallPapers) (*mtproto.Account_WallPapers, error) {
	c.Logger.Errorf("AccountGetWallPapers - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
