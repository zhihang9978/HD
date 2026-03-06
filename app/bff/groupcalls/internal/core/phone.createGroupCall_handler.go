package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// PhoneCreateGroupCall
func (c *GroupCallsCore) PhoneCreateGroupCall(in *mtproto.TLPhoneCreateGroupCall) (*mtproto.Updates, error) {
	c.Logger.Errorf("PhoneCreateGroupCall - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
