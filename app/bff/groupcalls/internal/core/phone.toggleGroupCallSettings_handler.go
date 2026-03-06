package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// PhoneToggleGroupCallSettings
func (c *GroupCallsCore) PhoneToggleGroupCallSettings(in *mtproto.TLPhoneToggleGroupCallSettings) (*mtproto.Updates, error) {
	c.Logger.Errorf("PhoneToggleGroupCallSettings - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
