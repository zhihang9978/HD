package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesSearchStickers
func (c *StickersCore) MessagesSearchStickers(in *mtproto.TLMessagesSearchStickers) (*mtproto.Messages_FoundStickers, error) {
	c.Logger.Errorf("MessagesSearchStickers - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
