package core

import "github.com/teamgram/proto/mtproto"

func (c *NotificationCore) AccountRegisterDevice(in *mtproto.TLAccountRegisterDevice) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
