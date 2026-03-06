package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsReadMessageContents
func (c *ChannelsCore) ChannelsReadMessageContents(in *mtproto.TLChannelsReadMessageContents) (*mtproto.Bool, error) {
	c.Logger.Errorf("ChannelsReadMessageContents - not impl")

	return mtproto.BoolTrue, nil
}
