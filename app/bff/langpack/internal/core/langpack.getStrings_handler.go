package core

import (
	"github.com/teamgram/proto/mtproto"
)

// LangpackGetStrings
func (c *LangpackCore) LangpackGetStrings(in *mtproto.TLLangpackGetStrings) (*mtproto.Vector_LangPackString, error) {
	c.Logger.Errorf("LangpackGetStrings - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
