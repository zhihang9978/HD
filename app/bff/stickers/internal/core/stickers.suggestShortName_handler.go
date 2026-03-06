package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StickersSuggestShortName
func (c *StickersCore) StickersSuggestShortName(in *mtproto.TLStickersSuggestShortName) (*mtproto.Stickers_SuggestedShortName, error) {
	c.Logger.Errorf("StickersSuggestShortName - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
