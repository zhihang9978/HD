package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PaymentsGetBankCardData
func (c *PaymentsCore) PaymentsGetBankCardData(in *mtproto.TLPaymentsGetBankCardData) (*mtproto.Payments_BankCardData, error) {
	c.Logger.Errorf("PaymentsGetBankCardData - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
