package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesSearchEmojiStickerSets
func (c *StickersCore) MessagesSearchEmojiStickerSets(in *mtproto.TLMessagesSearchEmojiStickerSets) (*mtproto.Messages_FoundStickerSets, error) {
	c.Logger.Errorf("MessagesSearchEmojiStickerSets - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
