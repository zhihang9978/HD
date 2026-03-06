package core

import "github.com/teamgram/proto/mtproto"

func (c *PremiumCore) HelpGetPremiumPromo(in *mtproto.TLHelpGetPremiumPromo) (*mtproto.Help_PremiumPromo, error) {
	return mtproto.MakeTLHelpPremiumPromo(&mtproto.Help_PremiumPromo{
		StatusText: "Premium is not available", StatusEntities: []*mtproto.MessageEntity{},
		VideoSections: []string{}, Videos: []*mtproto.Document{},
		PeriodOptions: []*mtproto.PremiumSubscriptionOption{}, Users: []*mtproto.User{},
		Currency: "USD", MonthlyAmount: 0,
	}).To_Help_PremiumPromo(), nil
}
