package core

import "github.com/teamgram/proto/mtproto"

func (c *ContactsCore) ContactsUpdateContactNote(in *mtproto.TLContactsUpdateContactNote) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
