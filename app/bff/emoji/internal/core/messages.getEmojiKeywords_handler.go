package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetEmojiKeywords
func (c *EmojiCore) MessagesGetEmojiKeywords(in *mtproto.TLMessagesGetEmojiKeywords) (*mtproto.EmojiKeywordsDifference, error) {
	c.Logger.Errorf("MessagesGetEmojiKeywords - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
