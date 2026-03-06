package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesDeletePhoneCallHistory
func (c *VoipCallsCore) MessagesDeletePhoneCallHistory(in *mtproto.TLMessagesDeletePhoneCallHistory) (*mtproto.Messages_AffectedFoundMessages, error) {
	c.Logger.Errorf("MessagesDeletePhoneCallHistory - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
