package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StickersChangeStickerPosition
func (c *StickersCore) StickersChangeStickerPosition(in *mtproto.TLStickersChangeStickerPosition) (*mtproto.Messages_StickerSet, error) {
	c.Logger.Errorf("StickersChangeStickerPosition - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
