package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesClearRecentReactions
func (c *ReactionsCore) MessagesClearRecentReactions(in *mtproto.TLMessagesClearRecentReactions) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesClearRecentReactions - not impl")

	return mtproto.BoolTrue, nil
}
