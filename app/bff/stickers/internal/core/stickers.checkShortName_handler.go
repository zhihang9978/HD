package core

import (
	"github.com/teamgram/proto/mtproto"
)

// StickersCheckShortName
func (c *StickersCore) StickersCheckShortName(in *mtproto.TLStickersCheckShortName) (*mtproto.Bool, error) {
	c.Logger.Errorf("StickersCheckShortName - not impl")

	return mtproto.BoolTrue, nil
}
