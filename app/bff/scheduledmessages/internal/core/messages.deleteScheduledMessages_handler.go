package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// MessagesDeleteScheduledMessages
func (c *ScheduledMessagesCore) MessagesDeleteScheduledMessages(in *mtproto.TLMessagesDeleteScheduledMessages) (*mtproto.Updates, error) {
	c.Logger.Errorf("MessagesDeleteScheduledMessages - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
