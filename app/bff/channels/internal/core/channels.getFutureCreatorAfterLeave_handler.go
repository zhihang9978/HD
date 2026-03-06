package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsGetFutureCreatorAfterLeave
func (c *ChannelsCore) ChannelsGetFutureCreatorAfterLeave(in *mtproto.TLChannelsGetFutureCreatorAfterLeave) (*mtproto.User, error) {
	c.Logger.Errorf("ChannelsGetFutureCreatorAfterLeave - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
