package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesReadFeaturedStickers
func (c *StickersCore) MessagesReadFeaturedStickers(in *mtproto.TLMessagesReadFeaturedStickers) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesReadFeaturedStickers - not impl")

	return mtproto.BoolTrue, nil
}
