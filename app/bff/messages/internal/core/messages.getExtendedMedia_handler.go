package core

import "github.com/teamgram/proto/mtproto"

func (c *MessagesCore) MessagesGetExtendedMedia(in *mtproto.TLMessagesGetExtendedMedia) (*mtproto.Updates, error) {
	return mtproto.MakeTLUpdates(&mtproto.Updates{Updates: []*mtproto.Update{}, Users: []*mtproto.User{}, Chats: []*mtproto.Chat{}}).To_Updates(), nil
}
