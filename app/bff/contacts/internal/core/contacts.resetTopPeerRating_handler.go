package core

import "github.com/teamgram/proto/mtproto"

func (c *ContactsCore) ContactsResetTopPeerRating(in *mtproto.TLContactsResetTopPeerRating) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
