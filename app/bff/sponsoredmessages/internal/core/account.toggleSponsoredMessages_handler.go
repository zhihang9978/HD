package core

import "github.com/teamgram/proto/mtproto"

func (c *SponsoredMessagesCore) AccountToggleSponsoredMessages(in *mtproto.TLAccountToggleSponsoredMessages) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
