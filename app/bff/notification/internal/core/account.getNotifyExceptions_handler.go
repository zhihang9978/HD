package core

import "github.com/teamgram/proto/mtproto"

func (c *NotificationCore) AccountGetNotifyExceptions(in *mtproto.TLAccountGetNotifyExceptions) (*mtproto.Updates, error) {
	return mtproto.MakeTLUpdates(&mtproto.Updates{Updates: []*mtproto.Update{}, Users: []*mtproto.User{}, Chats: []*mtproto.Chat{}}).To_Updates(), nil
}
