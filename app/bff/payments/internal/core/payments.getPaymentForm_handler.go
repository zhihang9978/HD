package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PaymentsGetPaymentForm
func (c *PaymentsCore) PaymentsGetPaymentForm(in *mtproto.TLPaymentsGetPaymentForm) (*mtproto.Payments_PaymentForm, error) {
	c.Logger.Errorf("PaymentsGetPaymentForm - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
