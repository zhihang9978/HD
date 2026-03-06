package core

import "github.com/teamgram/proto/mtproto"

func (c *MessagesCore) MessagesGetMessageEditData(in *mtproto.TLMessagesGetMessageEditData) (*mtproto.Messages_MessageEditData, error) {
	return mtproto.MakeTLMessagesMessageEditData(&mtproto.Messages_MessageEditData{Caption: false}).To_Messages_MessageEditData(), nil
}
