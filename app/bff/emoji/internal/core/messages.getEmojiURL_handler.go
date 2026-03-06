package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetEmojiURL
func (c *EmojiCore) MessagesGetEmojiURL(in *mtproto.TLMessagesGetEmojiURL) (*mtproto.EmojiURL, error) {
	c.Logger.Errorf("MessagesGetEmojiURL - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
