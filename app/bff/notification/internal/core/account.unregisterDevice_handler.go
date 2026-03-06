package core

import "github.com/teamgram/proto/mtproto"

func (c *NotificationCore) AccountUnregisterDevice(in *mtproto.TLAccountUnregisterDevice) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
