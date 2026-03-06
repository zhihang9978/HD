package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// PhoneEditGroupCallTitle
func (c *GroupCallsCore) PhoneEditGroupCallTitle(in *mtproto.TLPhoneEditGroupCallTitle) (*mtproto.Updates, error) {
	c.Logger.Errorf("PhoneEditGroupCallTitle - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
