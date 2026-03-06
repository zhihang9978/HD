package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesSearchStickerSets
func (c *StickersCore) MessagesSearchStickerSets(in *mtproto.TLMessagesSearchStickerSets) (*mtproto.Messages_FoundStickerSets, error) {
	c.Logger.Errorf("MessagesSearchStickerSets - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
