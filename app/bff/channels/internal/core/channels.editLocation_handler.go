package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsEditLocation
func (c *ChannelsCore) ChannelsEditLocation(in *mtproto.TLChannelsEditLocation) (*mtproto.Bool, error) {
	c.Logger.Errorf("ChannelsEditLocation - not impl")

	return mtproto.BoolTrue, nil
}
