package core

import "github.com/teamgram/proto/mtproto"

func (c *AccountCore) AccountConfirmPhone(in *mtproto.TLAccountConfirmPhone) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
