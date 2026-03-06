package core

import "github.com/teamgram/proto/mtproto"

func (c *ContactsCore) ContactsToggleTopPeers(in *mtproto.TLContactsToggleTopPeers) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
