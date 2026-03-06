package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// PhoneJoinGroupCall
func (c *GroupCallsCore) PhoneJoinGroupCall(in *mtproto.TLPhoneJoinGroupCall) (*mtproto.Updates, error) {
	c.Logger.Errorf("PhoneJoinGroupCall - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
