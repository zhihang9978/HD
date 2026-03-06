package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// PhoneDeleteGroupCallParticipantMessages
func (c *GroupCallsCore) PhoneDeleteGroupCallParticipantMessages(in *mtproto.TLPhoneDeleteGroupCallParticipantMessages) (*mtproto.Updates, error) {
	c.Logger.Errorf("PhoneDeleteGroupCallParticipantMessages - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
