package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StickersSetStickerSetThumb
func (c *StickersCore) StickersSetStickerSetThumb(in *mtproto.TLStickersSetStickerSetThumb) (*mtproto.Messages_StickerSet, error) {
	c.Logger.Errorf("StickersSetStickerSetThumb - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
