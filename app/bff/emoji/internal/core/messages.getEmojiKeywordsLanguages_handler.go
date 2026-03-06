package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetEmojiKeywordsLanguages
func (c *EmojiCore) MessagesGetEmojiKeywordsLanguages(in *mtproto.TLMessagesGetEmojiKeywordsLanguages) (*mtproto.Vector_EmojiLanguage, error) {
	c.Logger.Errorf("MessagesGetEmojiKeywordsLanguages - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
