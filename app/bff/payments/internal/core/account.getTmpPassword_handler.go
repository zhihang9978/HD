package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountGetTmpPassword
func (c *PaymentsCore) AccountGetTmpPassword(in *mtproto.TLAccountGetTmpPassword) (*mtproto.Account_TmpPassword, error) {
	c.Logger.Errorf("AccountGetTmpPassword - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
