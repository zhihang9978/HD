package core

import "github.com/teamgram/proto/mtproto"

func (c *ContactsCore) ContactsResetSaved(in *mtproto.TLContactsResetSaved) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
