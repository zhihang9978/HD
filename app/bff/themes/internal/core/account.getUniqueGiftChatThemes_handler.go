package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountGetUniqueGiftChatThemes
func (c *ThemesCore) AccountGetUniqueGiftChatThemes(in *mtproto.TLAccountGetUniqueGiftChatThemes) (*mtproto.Account_ChatThemes, error) {
	c.Logger.Errorf("AccountGetUniqueGiftChatThemes - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
