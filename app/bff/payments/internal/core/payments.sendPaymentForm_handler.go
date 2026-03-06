package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PaymentsSendPaymentForm
func (c *PaymentsCore) PaymentsSendPaymentForm(in *mtproto.TLPaymentsSendPaymentForm) (*mtproto.Payments_PaymentResult, error) {
	c.Logger.Errorf("PaymentsSendPaymentForm - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
