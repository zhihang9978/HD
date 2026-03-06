package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesToggleStickerSets
func (c *StickersCore) MessagesToggleStickerSets(in *mtproto.TLMessagesToggleStickerSets) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesToggleStickerSets - not impl")

	return mtproto.BoolTrue, nil
}
