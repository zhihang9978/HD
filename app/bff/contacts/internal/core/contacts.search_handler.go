package core

import "github.com/teamgram/proto/mtproto"

func (c *ContactsCore) ContactsSearch(in *mtproto.TLContactsSearch) (*mtproto.Contacts_Found, error) {
	return mtproto.MakeTLContactsFound(&mtproto.Contacts_Found{
		MyResults: []*mtproto.Peer{}, Results: []*mtproto.Peer{}, Chats: []*mtproto.Chat{}, Users: []*mtproto.User{},
	}).To_Contacts_Found(), nil
}
