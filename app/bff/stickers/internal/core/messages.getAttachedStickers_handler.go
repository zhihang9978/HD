package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetAttachedStickers
func (c *StickersCore) MessagesGetAttachedStickers(in *mtproto.TLMessagesGetAttachedStickers) (*mtproto.Vector_StickerSetCovered, error) {
	c.Logger.Errorf("MessagesGetAttachedStickers - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
