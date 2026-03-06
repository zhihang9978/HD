package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetArchivedStickers
func (c *StickersCore) MessagesGetArchivedStickers(in *mtproto.TLMessagesGetArchivedStickers) (*mtproto.Messages_ArchivedStickers, error) {
	c.Logger.Errorf("MessagesGetArchivedStickers - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
