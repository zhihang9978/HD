package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesAcceptUrlAuth
func (c *SeamlessCore) MessagesAcceptUrlAuth(in *mtproto.TLMessagesAcceptUrlAuth) (*mtproto.UrlAuthResult, error) {
	c.Logger.Errorf("MessagesAcceptUrlAuth - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
