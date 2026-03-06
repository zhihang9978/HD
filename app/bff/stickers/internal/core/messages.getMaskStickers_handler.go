package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetMaskStickers
func (c *StickersCore) MessagesGetMaskStickers(in *mtproto.TLMessagesGetMaskStickers) (*mtproto.Messages_AllStickers, error) {
	c.Logger.Errorf("MessagesGetMaskStickers - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
