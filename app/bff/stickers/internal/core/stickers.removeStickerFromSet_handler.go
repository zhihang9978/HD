package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StickersRemoveStickerFromSet
func (c *StickersCore) StickersRemoveStickerFromSet(in *mtproto.TLStickersRemoveStickerFromSet) (*mtproto.Messages_StickerSet, error) {
	c.Logger.Errorf("StickersRemoveStickerFromSet - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
