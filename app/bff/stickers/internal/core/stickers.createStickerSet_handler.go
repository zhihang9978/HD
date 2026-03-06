package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StickersCreateStickerSet
func (c *StickersCore) StickersCreateStickerSet(in *mtproto.TLStickersCreateStickerSet) (*mtproto.Messages_StickerSet, error) {
	c.Logger.Errorf("StickersCreateStickerSet - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
