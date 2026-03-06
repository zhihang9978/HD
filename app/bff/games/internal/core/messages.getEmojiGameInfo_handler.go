package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetEmojiGameInfo
func (c *GamesCore) MessagesGetEmojiGameInfo(in *mtproto.TLMessagesGetEmojiGameInfo) (*mtproto.Messages_EmojiGameInfo, error) {
	c.Logger.Errorf("MessagesGetEmojiGameInfo - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
