package core

import (
	"github.com/teamgram/proto/mtproto"
)

// LangpackGetLanguage
func (c *LangpackCore) LangpackGetLanguage(in *mtproto.TLLangpackGetLanguage) (*mtproto.LangPackLanguage, error) {
	c.Logger.Errorf("LangpackGetLanguage - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
