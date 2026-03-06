package core

import "github.com/teamgram/proto/mtproto"

func (c *MessagesCore) MessagesSearchSentMedia(in *mtproto.TLMessagesSearchSentMedia) (*mtproto.Messages_Messages, error) {
	return mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{Messages: []*mtproto.Message{}, Chats: []*mtproto.Chat{}, Users: []*mtproto.User{}}).To_Messages_Messages(), nil
}
