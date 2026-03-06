package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// PhoneEditGroupCallParticipant
func (c *GroupCallsCore) PhoneEditGroupCallParticipant(in *mtproto.TLPhoneEditGroupCallParticipant) (*mtproto.Updates, error) {
	c.Logger.Errorf("PhoneEditGroupCallParticipant - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
