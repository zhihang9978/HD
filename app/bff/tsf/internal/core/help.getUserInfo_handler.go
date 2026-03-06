package core

import (
	"github.com/teamgram/proto/mtproto"
)

// HelpGetUserInfo
func (c *TsfCore) HelpGetUserInfo(in *mtproto.TLHelpGetUserInfo) (*mtproto.Help_UserInfo, error) {
	c.Logger.Errorf("HelpGetUserInfo - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
