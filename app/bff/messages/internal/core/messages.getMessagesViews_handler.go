package core

import "github.com/teamgram/proto/mtproto"

func (c *MessagesCore) MessagesGetMessagesViews(in *mtproto.TLMessagesGetMessagesViews) (*mtproto.Messages_MessageViews, error) {
	views := make([]*mtproto.MessageViews, 0, len(in.GetId()))
	for range in.GetId() {
		views = append(views, mtproto.MakeTLMessageViews(&mtproto.MessageViews{}).To_MessageViews())
	}
	return mtproto.MakeTLMessagesMessageViews(&mtproto.Messages_MessageViews{Views: views, Chats: []*mtproto.Chat{}, Users: []*mtproto.User{}}).To_Messages_MessageViews(), nil
}
