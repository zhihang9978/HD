package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PaymentsValidateRequestedInfo
func (c *PaymentsCore) PaymentsValidateRequestedInfo(in *mtproto.TLPaymentsValidateRequestedInfo) (*mtproto.Payments_ValidatedRequestedInfo, error) {
	c.Logger.Errorf("PaymentsValidateRequestedInfo - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
