package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetGameHighScores
func (c *GamesCore) MessagesGetGameHighScores(in *mtproto.TLMessagesGetGameHighScores) (*mtproto.Messages_HighScores, error) {
	c.Logger.Errorf("MessagesGetGameHighScores - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
