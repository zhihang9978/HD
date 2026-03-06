package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsReadHistory
func (c *ChannelsCore) ChannelsReadHistory(in *mtproto.TLChannelsReadHistory) (*mtproto.Bool, error) {
	c.Logger.Errorf("ChannelsReadHistory - not impl")

	return mtproto.BoolTrue, nil
}
