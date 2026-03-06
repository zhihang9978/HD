package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StickersRenameStickerSet
func (c *StickersCore) StickersRenameStickerSet(in *mtproto.TLStickersRenameStickerSet) (*mtproto.Messages_StickerSet, error) {
	c.Logger.Errorf("StickersRenameStickerSet - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
