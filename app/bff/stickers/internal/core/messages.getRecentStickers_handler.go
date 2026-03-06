package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetRecentStickers
func (c *StickersCore) MessagesGetRecentStickers(in *mtproto.TLMessagesGetRecentStickers) (*mtproto.Messages_RecentStickers, error) {
	c.Logger.Errorf("MessagesGetRecentStickers - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
