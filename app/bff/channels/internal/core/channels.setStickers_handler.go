package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsSetStickers
func (c *ChannelsCore) ChannelsSetStickers(in *mtproto.TLChannelsSetStickers) (*mtproto.Bool, error) {
	c.Logger.Errorf("ChannelsSetStickers - not impl")

	return mtproto.BoolTrue, nil
}
