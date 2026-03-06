package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetEmojiKeywordsDifference
func (c *EmojiCore) MessagesGetEmojiKeywordsDifference(in *mtproto.TLMessagesGetEmojiKeywordsDifference) (*mtproto.EmojiKeywordsDifference, error) {
	c.Logger.Errorf("MessagesGetEmojiKeywordsDifference - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
