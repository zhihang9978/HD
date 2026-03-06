package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// PhoneToggleGroupCallStartSubscription
func (c *GroupCallsCore) PhoneToggleGroupCallStartSubscription(in *mtproto.TLPhoneToggleGroupCallStartSubscription) (*mtproto.Updates, error) {
	c.Logger.Errorf("PhoneToggleGroupCallStartSubscription - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
