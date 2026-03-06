package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesFaveSticker
func (c *StickersCore) MessagesFaveSticker(in *mtproto.TLMessagesFaveSticker) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesFaveSticker - not impl")

	return mtproto.BoolTrue, nil
}
