package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsGetGroupsForDiscussion
func (c *ChannelsCore) ChannelsGetGroupsForDiscussion(in *mtproto.TLChannelsGetGroupsForDiscussion) (*mtproto.Messages_Chats, error) {
	c.Logger.Errorf("ChannelsGetGroupsForDiscussion - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
