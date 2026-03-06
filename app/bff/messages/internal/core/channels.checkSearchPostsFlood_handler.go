package core

import "github.com/teamgram/proto/mtproto"

func (c *MessagesCore) ChannelsCheckSearchPostsFlood(in *mtproto.TLChannelsCheckSearchPostsFlood) (*mtproto.SearchPostsFlood, error) {
	return mtproto.MakeTLSearchPostsFlood(&mtproto.SearchPostsFlood{}).To_SearchPostsFlood(), nil
}
