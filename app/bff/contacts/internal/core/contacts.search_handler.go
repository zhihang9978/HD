package core

import (
	"github.com/teamgram/proto/mtproto"
	userpb "github.com/teamgram/teamgram-server/app/service/biz/user/user"
)

func (c *ContactsCore) ContactsSearch(in *mtproto.TLContactsSearch) (*mtproto.Contacts_Found, error) {
	if in.Q == "" {
		return mtproto.MakeTLContactsFound(&mtproto.Contacts_Found{
			MyResults: []*mtproto.Peer{},
			Results:   []*mtproto.Peer{},
			Chats:     []*mtproto.Chat{},
			Users:     []*mtproto.User{},
		}).To_Contacts_Found(), nil
	}

	limit := in.Limit
	if limit <= 0 {
		limit = 50
	}
	if limit > 50 {
		limit = 50
	}

	// Search for users matching the query
	found, err := c.svcCtx.Dao.UserClient.UserSearch(c.ctx, &userpb.TLUserSearch{
		Q:                in.Q,
		ExcludedContacts: []int64{c.MD.UserId},
		Offset:           0,
		Limit:            limit,
	})
	if err != nil {
		c.Logger.Errorf("contacts.search - UserSearch error: %v", err)
		return mtproto.MakeTLContactsFound(&mtproto.Contacts_Found{
			MyResults: []*mtproto.Peer{},
			Results:   []*mtproto.Peer{},
			Chats:     []*mtproto.Chat{},
			Users:     []*mtproto.User{},
		}).To_Contacts_Found(), nil
	}

	if len(found.GetIdList()) == 0 {
		return mtproto.MakeTLContactsFound(&mtproto.Contacts_Found{
			MyResults: []*mtproto.Peer{},
			Results:   []*mtproto.Peer{},
			Chats:     []*mtproto.Chat{},
			Users:     []*mtproto.User{},
		}).To_Contacts_Found(), nil
	}

	// Get full user objects
	idList := found.GetIdList()
	idList = append(idList, c.MD.UserId) // include self for ToUnsafeUser context

	mUsers, err := c.svcCtx.Dao.UserClient.UserGetMutableUsers(c.ctx, &userpb.TLUserGetMutableUsers{
		Id: idList,
	})
	if err != nil {
		c.Logger.Errorf("contacts.search - UserGetMutableUsers error: %v", err)
		return mtproto.MakeTLContactsFound(&mtproto.Contacts_Found{
			MyResults: []*mtproto.Peer{},
			Results:   []*mtproto.Peer{},
			Chats:     []*mtproto.Chat{},
			Users:     []*mtproto.User{},
		}).To_Contacts_Found(), nil
	}

	// Build results
	results := make([]*mtproto.Peer, 0, len(found.GetIdList()))
	for _, uid := range found.GetIdList() {
		if uid != c.MD.UserId {
			results = append(results, mtproto.MakeTLPeerUser(&mtproto.Peer{UserId: uid}).To_Peer())
		}
	}

	users := mUsers.GetUserListByIdList(c.MD.UserId, found.GetIdList()...)

	return mtproto.MakeTLContactsFound(&mtproto.Contacts_Found{
		MyResults: []*mtproto.Peer{},
		Results:   results,
		Chats:     []*mtproto.Chat{},
		Users:     users,
	}).To_Contacts_Found(), nil
}
