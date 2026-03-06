package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// ContactsBlockFromReplies
func (c *MessageThreadsCore) ContactsBlockFromReplies(in *mtproto.TLContactsBlockFromReplies) (*mtproto.Updates, error) {
	c.Logger.Errorf("ContactsBlockFromReplies - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
