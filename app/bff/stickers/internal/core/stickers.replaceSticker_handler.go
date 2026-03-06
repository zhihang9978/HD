package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StickersReplaceSticker
func (c *StickersCore) StickersReplaceSticker(in *mtproto.TLStickersReplaceSticker) (*mtproto.Messages_StickerSet, error) {
	c.Logger.Errorf("StickersReplaceSticker - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
