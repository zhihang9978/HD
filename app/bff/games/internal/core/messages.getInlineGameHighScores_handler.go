package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetInlineGameHighScores
func (c *GamesCore) MessagesGetInlineGameHighScores(in *mtproto.TLMessagesGetInlineGameHighScores) (*mtproto.Messages_HighScores, error) {
	c.Logger.Errorf("MessagesGetInlineGameHighScores - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
