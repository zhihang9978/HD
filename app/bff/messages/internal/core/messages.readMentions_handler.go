package core

import "github.com/teamgram/proto/mtproto"

func (c *MessagesCore) MessagesReadMentions(in *mtproto.TLMessagesReadMentions) (*mtproto.Messages_AffectedHistory, error) {
	return mtproto.MakeTLMessagesAffectedHistory(&mtproto.Messages_AffectedHistory{Pts: 0, PtsCount: 0, Offset: 0}).To_Messages_AffectedHistory(), nil
}
