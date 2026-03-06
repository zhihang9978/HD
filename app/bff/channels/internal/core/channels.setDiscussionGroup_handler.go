package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsSetDiscussionGroup
func (c *ChannelsCore) ChannelsSetDiscussionGroup(in *mtproto.TLChannelsSetDiscussionGroup) (*mtproto.Bool, error) {
	c.Logger.Errorf("ChannelsSetDiscussionGroup - not impl")

	return mtproto.BoolTrue, nil
}
