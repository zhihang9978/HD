package core

import "github.com/teamgram/proto/mtproto"

func (c *ContactsCore) ContactsSetBlocked(in *mtproto.TLContactsSetBlocked) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
