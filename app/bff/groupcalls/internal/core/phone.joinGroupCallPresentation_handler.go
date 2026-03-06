package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// PhoneJoinGroupCallPresentation
func (c *GroupCallsCore) PhoneJoinGroupCallPresentation(in *mtproto.TLPhoneJoinGroupCallPresentation) (*mtproto.Updates, error) {
	c.Logger.Errorf("PhoneJoinGroupCallPresentation - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
