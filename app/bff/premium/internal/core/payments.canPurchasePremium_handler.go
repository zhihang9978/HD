package core

import "github.com/teamgram/proto/mtproto"

func (c *PremiumCore) PaymentsCanPurchasePremium(in *mtproto.TLPaymentsCanPurchasePremium) (*mtproto.Bool, error) {
	return mtproto.BoolFalse, nil
}
