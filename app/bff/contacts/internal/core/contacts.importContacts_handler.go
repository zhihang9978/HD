package core

import "github.com/teamgram/proto/mtproto"

func (c *ContactsCore) ContactsImportContacts(in *mtproto.TLContactsImportContacts) (*mtproto.Contacts_ImportedContacts, error) {
	return mtproto.MakeTLContactsImportedContacts(&mtproto.Contacts_ImportedContacts{
		Imported: []*mtproto.ImportedContact{}, PopularInvites: []*mtproto.PopularContact{},
		RetryContacts: []int64{}, Users: []*mtproto.User{},
	}).To_Contacts_ImportedContacts(), nil
}
