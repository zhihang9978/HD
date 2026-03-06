package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneExportGroupCallInvite
func (c *GroupCallsCore) PhoneExportGroupCallInvite(in *mtproto.TLPhoneExportGroupCallInvite) (*mtproto.Phone_ExportedGroupCallInvite, error) {
	c.Logger.Errorf("PhoneExportGroupCallInvite - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
