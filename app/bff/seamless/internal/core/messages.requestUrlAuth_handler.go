package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesRequestUrlAuth
func (c *SeamlessCore) MessagesRequestUrlAuth(in *mtproto.TLMessagesRequestUrlAuth) (*mtproto.UrlAuthResult, error) {
	c.Logger.Errorf("MessagesRequestUrlAuth - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
