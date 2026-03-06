package core

import (
	"github.com/teamgram/proto/mtproto"
)

// HelpEditUserInfo
func (c *TsfCore) HelpEditUserInfo(in *mtproto.TLHelpEditUserInfo) (*mtproto.Help_UserInfo, error) {
	c.Logger.Errorf("HelpEditUserInfo - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
