package core

import "github.com/teamgram/proto/mtproto"

func (c *NotificationCore) AccountUpdateNotifySettings(in *mtproto.TLAccountUpdateNotifySettings) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
