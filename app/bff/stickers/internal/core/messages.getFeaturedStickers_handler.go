package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetFeaturedStickers
func (c *StickersCore) MessagesGetFeaturedStickers(in *mtproto.TLMessagesGetFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error) {
	c.Logger.Errorf("MessagesGetFeaturedStickers - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
