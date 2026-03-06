package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesClearRecentStickers
func (c *StickersCore) MessagesClearRecentStickers(in *mtproto.TLMessagesClearRecentStickers) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesClearRecentStickers - not impl")

	return mtproto.BoolTrue, nil
}
