package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PhoneGetGroupParticipants
func (c *GroupCallsCore) PhoneGetGroupParticipants(in *mtproto.TLPhoneGetGroupParticipants) (*mtproto.Phone_GroupParticipants, error) {
	c.Logger.Errorf("PhoneGetGroupParticipants - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
