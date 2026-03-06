package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// PhoneDiscardGroupCall
func (c *GroupCallsCore) PhoneDiscardGroupCall(in *mtproto.TLPhoneDiscardGroupCall) (*mtproto.Updates, error) {
	c.Logger.Errorf("PhoneDiscardGroupCall - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
