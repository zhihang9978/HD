package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StickersDeleteStickerSet
func (c *StickersCore) StickersDeleteStickerSet(in *mtproto.TLStickersDeleteStickerSet) (*mtproto.Bool, error) {
	c.Logger.Errorf("StickersDeleteStickerSet - not impl")

	return mtproto.BoolTrue, nil
}
