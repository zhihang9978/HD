package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesSetDefaultReaction
func (c *ReactionsCore) MessagesSetDefaultReaction(in *mtproto.TLMessagesSetDefaultReaction) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesSetDefaultReaction - not impl")

	return mtproto.BoolTrue, nil
}
