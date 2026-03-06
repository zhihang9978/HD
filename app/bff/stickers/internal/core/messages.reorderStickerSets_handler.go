package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesReorderStickerSets
func (c *StickersCore) MessagesReorderStickerSets(in *mtproto.TLMessagesReorderStickerSets) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesReorderStickerSets - not impl")

	return mtproto.BoolTrue, nil
}
