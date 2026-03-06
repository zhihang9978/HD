package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PaymentsGetPaymentReceipt
func (c *PaymentsCore) PaymentsGetPaymentReceipt(in *mtproto.TLPaymentsGetPaymentReceipt) (*mtproto.Payments_PaymentReceipt, error) {
	c.Logger.Errorf("PaymentsGetPaymentReceipt - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
