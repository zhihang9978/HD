package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetAllStickers
func (c *StickersCore) MessagesGetAllStickers(in *mtproto.TLMessagesGetAllStickers) (*mtproto.Messages_AllStickers, error) {
	c.Logger.Errorf("MessagesGetAllStickers - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
