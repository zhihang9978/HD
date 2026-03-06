package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// MessagesSendScheduledMessages
func (c *ScheduledMessagesCore) MessagesSendScheduledMessages(in *mtproto.TLMessagesSendScheduledMessages) (*mtproto.Updates, error) {
	c.Logger.Errorf("MessagesSendScheduledMessages - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
