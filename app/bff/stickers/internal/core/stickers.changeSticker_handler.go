package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StickersChangeSticker
func (c *StickersCore) StickersChangeSticker(in *mtproto.TLStickersChangeSticker) (*mtproto.Messages_StickerSet, error) {
	c.Logger.Errorf("StickersChangeSticker - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
