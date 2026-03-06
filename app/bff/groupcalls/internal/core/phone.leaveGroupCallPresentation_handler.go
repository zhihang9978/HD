package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// PhoneLeaveGroupCallPresentation
func (c *GroupCallsCore) PhoneLeaveGroupCallPresentation(in *mtproto.TLPhoneLeaveGroupCallPresentation) (*mtproto.Updates, error) {
	c.Logger.Errorf("PhoneLeaveGroupCallPresentation - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
