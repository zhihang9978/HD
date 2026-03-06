package core

import (
	"github.com/teamgram/proto/mtproto"
)

// HelpGetRecentMeUrls
func (c *DeepLinksCore) HelpGetRecentMeUrls(in *mtproto.TLHelpGetRecentMeUrls) (*mtproto.Help_RecentMeUrls, error) {
	c.Logger.Errorf("HelpGetRecentMeUrls - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
