package core

import "github.com/teamgram/proto/mtproto"

func (c *MessagesCore) MessagesSaveDefaultSendAs(in *mtproto.TLMessagesSaveDefaultSendAs) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
