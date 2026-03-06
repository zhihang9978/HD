package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountGetMultiWallPapers
func (c *WallpapersCore) AccountGetMultiWallPapers(in *mtproto.TLAccountGetMultiWallPapers) (*mtproto.Vector_WallPaper, error) {
	c.Logger.Errorf("AccountGetMultiWallPapers - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
