package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetFavedStickers
func (c *StickersCore) MessagesGetFavedStickers(in *mtproto.TLMessagesGetFavedStickers) (*mtproto.Messages_FavedStickers, error) {
	c.Logger.Errorf("MessagesGetFavedStickers - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
