package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// MessagesSendVote
func (c *PollsCore) MessagesSendVote(in *mtproto.TLMessagesSendVote) (*mtproto.Updates, error) {
	c.Logger.Errorf("MessagesSendVote - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
