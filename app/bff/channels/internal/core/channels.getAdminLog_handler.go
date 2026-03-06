package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsGetAdminLog
func (c *ChannelsCore) ChannelsGetAdminLog(in *mtproto.TLChannelsGetAdminLog) (*mtproto.Channels_AdminLogResults, error) {
	c.Logger.Errorf("ChannelsGetAdminLog - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
