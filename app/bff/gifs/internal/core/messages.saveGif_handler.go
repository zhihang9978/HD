package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesSaveGif
func (c *GifsCore) MessagesSaveGif(in *mtproto.TLMessagesSaveGif) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesSaveGif - not impl")

	return mtproto.BoolTrue, nil
}
