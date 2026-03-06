package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PaymentsClearSavedInfo
func (c *PaymentsCore) PaymentsClearSavedInfo(in *mtproto.TLPaymentsClearSavedInfo) (*mtproto.Bool, error) {
	c.Logger.Errorf("PaymentsClearSavedInfo - not impl")

	return mtproto.BoolTrue, nil
}
