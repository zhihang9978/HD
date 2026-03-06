package core

import (
	"github.com/teamgram/proto/mtproto"
)

// LangpackGetDifference
func (c *LangpackCore) LangpackGetDifference(in *mtproto.TLLangpackGetDifference) (*mtproto.LangPackDifference, error) {
	c.Logger.Errorf("LangpackGetDifference - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
