package core

import "github.com/teamgram/proto/mtproto"

func (c *MessagesCore) MessagesGetSearchResultsCalendar(in *mtproto.TLMessagesGetSearchResultsCalendar) (*mtproto.Messages_SearchResultsCalendar, error) {
	return mtproto.MakeTLMessagesSearchResultsCalendar(&mtproto.Messages_SearchResultsCalendar{
		Count: 0, MinDate: 0, MinMsgId: 0,
		Periods: []*mtproto.SearchResultsCalendarPeriod{}, Messages: []*mtproto.Message{},
		Chats: []*mtproto.Chat{}, Users: []*mtproto.User{},
	}).To_Messages_SearchResultsCalendar(), nil
}
