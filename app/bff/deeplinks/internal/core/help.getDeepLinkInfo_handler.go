package core

import (
	"github.com/teamgram/proto/mtproto"
)

// HelpGetDeepLinkInfo
func (c *DeepLinksCore) HelpGetDeepLinkInfo(in *mtproto.TLHelpGetDeepLinkInfo) (*mtproto.Help_DeepLinkInfo, error) {
	c.Logger.Errorf("HelpGetDeepLinkInfo - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
