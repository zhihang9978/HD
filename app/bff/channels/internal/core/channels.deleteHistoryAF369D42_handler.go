package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsDeleteHistoryAF369D42
func (c *ChannelsCore) ChannelsDeleteHistoryAF369D42(in *mtproto.TLChannelsDeleteHistoryAF369D42) (*mtproto.Bool, error) {
	c.Logger.Errorf("ChannelsDeleteHistoryAF369D42 - not impl")

	return mtproto.BoolTrue, nil
}
