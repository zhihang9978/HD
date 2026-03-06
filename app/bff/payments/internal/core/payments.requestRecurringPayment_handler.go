package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// PaymentsRequestRecurringPayment
func (c *PaymentsCore) PaymentsRequestRecurringPayment(in *mtproto.TLPaymentsRequestRecurringPayment) (*mtproto.Updates, error) {
	c.Logger.Errorf("PaymentsRequestRecurringPayment - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
