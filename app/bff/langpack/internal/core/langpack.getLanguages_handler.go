package core

import (
	"github.com/teamgram/proto/mtproto"
)

// LangpackGetLanguages
func (c *LangpackCore) LangpackGetLanguages(in *mtproto.TLLangpackGetLanguages) (*mtproto.Vector_LangPackLanguage, error) {
	c.Logger.Errorf("LangpackGetLanguages - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
