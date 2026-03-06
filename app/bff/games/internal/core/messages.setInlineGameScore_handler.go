package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesSetInlineGameScore
func (c *GamesCore) MessagesSetInlineGameScore(in *mtproto.TLMessagesSetInlineGameScore) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesSetInlineGameScore - not impl")

	return mtproto.BoolTrue, nil
}
