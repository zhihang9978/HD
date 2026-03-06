package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetOldFeaturedStickers
func (c *StickersCore) MessagesGetOldFeaturedStickers(in *mtproto.TLMessagesGetOldFeaturedStickers) (*mtproto.Messages_FeaturedStickers, error) {
	c.Logger.Errorf("MessagesGetOldFeaturedStickers - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
