package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountResetWallPapers
func (c *WallpapersCore) AccountResetWallPapers(in *mtproto.TLAccountResetWallPapers) (*mtproto.Bool, error) {
	c.Logger.Errorf("AccountResetWallPapers - not impl")

	return mtproto.BoolTrue, nil
}
