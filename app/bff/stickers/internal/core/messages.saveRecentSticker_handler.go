package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesSaveRecentSticker
func (c *StickersCore) MessagesSaveRecentSticker(in *mtproto.TLMessagesSaveRecentSticker) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesSaveRecentSticker - not impl")

	return mtproto.BoolTrue, nil
}
