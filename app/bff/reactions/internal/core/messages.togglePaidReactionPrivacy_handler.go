package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesTogglePaidReactionPrivacy
func (c *ReactionsCore) MessagesTogglePaidReactionPrivacy(in *mtproto.TLMessagesTogglePaidReactionPrivacy) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesTogglePaidReactionPrivacy - not impl")

	return mtproto.BoolTrue, nil
}
