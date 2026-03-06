package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// PhoneDeleteGroupCallMessages
func (c *GroupCallsCore) PhoneDeleteGroupCallMessages(in *mtproto.TLPhoneDeleteGroupCallMessages) (*mtproto.Updates, error) {
	c.Logger.Errorf("PhoneDeleteGroupCallMessages - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
