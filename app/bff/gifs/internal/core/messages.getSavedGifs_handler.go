package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetSavedGifs
func (c *GifsCore) MessagesGetSavedGifs(in *mtproto.TLMessagesGetSavedGifs) (*mtproto.Messages_SavedGifs, error) {
	c.Logger.Errorf("MessagesGetSavedGifs - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
