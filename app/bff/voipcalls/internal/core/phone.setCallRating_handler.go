package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// PhoneSetCallRating
func (c *VoipCallsCore) PhoneSetCallRating(in *mtproto.TLPhoneSetCallRating) (*mtproto.Updates, error) {
	c.Logger.Errorf("PhoneSetCallRating - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
