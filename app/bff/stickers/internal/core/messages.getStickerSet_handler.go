package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetStickerSet
func (c *StickersCore) MessagesGetStickerSet(in *mtproto.TLMessagesGetStickerSet) (*mtproto.Messages_StickerSet, error) {
	c.Logger.Errorf("MessagesGetStickerSet - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
