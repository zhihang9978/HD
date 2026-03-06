package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// MessagesSetGameScore
func (c *GamesCore) MessagesSetGameScore(in *mtproto.TLMessagesSetGameScore) (*mtproto.Updates, error) {
	c.Logger.Errorf("MessagesSetGameScore - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
