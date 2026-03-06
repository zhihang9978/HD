package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// PhoneSendGroupCallMessageB1D11410
func (c *GroupCallsCore) PhoneSendGroupCallMessageB1D11410(in *mtproto.TLPhoneSendGroupCallMessageB1D11410) (*mtproto.Updates, error) {
	c.Logger.Errorf("PhoneSendGroupCallMessageB1D11410 - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
