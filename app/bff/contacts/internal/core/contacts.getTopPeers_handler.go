package core

import "github.com/teamgram/proto/mtproto"

func (c *ContactsCore) ContactsGetTopPeers(in *mtproto.TLContactsGetTopPeers) (*mtproto.Contacts_TopPeers, error) {
	return mtproto.MakeTLContactsTopPeers(&mtproto.Contacts_TopPeers{
		Categories: []*mtproto.TopPeerCategoryPeers{}, Chats: []*mtproto.Chat{}, Users: []*mtproto.User{},
	}).To_Contacts_TopPeers(), nil
}
