package core

import (
	"github.com/teamgram/proto/mtproto"
)

// LangpackGetLangPack
func (c *LangpackCore) LangpackGetLangPack(in *mtproto.TLLangpackGetLangPack) (*mtproto.LangPackDifference, error) {
	c.Logger.Errorf("LangpackGetLangPack - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
