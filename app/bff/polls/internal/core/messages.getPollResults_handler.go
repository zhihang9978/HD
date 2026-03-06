package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetPollResults
func (c *PollsCore) MessagesGetPollResults(in *mtproto.TLMessagesGetPollResults) (*mtproto.Updates, error) {
	c.Logger.Errorf("MessagesGetPollResults - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
