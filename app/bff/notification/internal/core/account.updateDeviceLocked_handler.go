package core

import "github.com/teamgram/proto/mtproto"

func (c *NotificationCore) AccountUpdateDeviceLocked(in *mtproto.TLAccountUpdateDeviceLocked) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
