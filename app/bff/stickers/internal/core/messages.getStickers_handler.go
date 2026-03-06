package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetStickers
func (c *StickersCore) MessagesGetStickers(in *mtproto.TLMessagesGetStickers) (*mtproto.Messages_Stickers, error) {
	c.Logger.Errorf("MessagesGetStickers - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
