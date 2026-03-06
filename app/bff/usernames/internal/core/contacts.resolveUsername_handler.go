package core

import "github.com/teamgram/proto/mtproto"

func (c *UsernamesCore) ContactsResolveUsername(in *mtproto.TLContactsResolveUsername) (*mtproto.Contacts_ResolvedPeer, error) {
	return nil, mtproto.ErrUsernameNotOccupied
}
