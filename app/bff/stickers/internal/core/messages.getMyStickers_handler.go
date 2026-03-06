package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetMyStickers
func (c *StickersCore) MessagesGetMyStickers(in *mtproto.TLMessagesGetMyStickers) (*mtproto.Messages_MyStickers, error) {
	c.Logger.Errorf("MessagesGetMyStickers - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
