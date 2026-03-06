package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StickersAddStickerToSet
func (c *StickersCore) StickersAddStickerToSet(in *mtproto.TLStickersAddStickerToSet) (*mtproto.Messages_StickerSet, error) {
	c.Logger.Errorf("StickersAddStickerToSet - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
