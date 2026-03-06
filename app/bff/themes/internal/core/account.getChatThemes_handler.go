package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountGetChatThemes
func (c *ThemesCore) AccountGetChatThemes(in *mtproto.TLAccountGetChatThemes) (*mtproto.Account_Themes, error) {
	c.Logger.Errorf("AccountGetChatThemes - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
