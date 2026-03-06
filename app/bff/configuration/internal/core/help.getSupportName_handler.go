package core

import "github.com/teamgram/proto/mtproto"

func (c *ConfigurationCore) HelpGetSupportName(in *mtproto.TLHelpGetSupportName) (*mtproto.Help_SupportName, error) {
	return mtproto.MakeTLHelpSupportName(&mtproto.Help_SupportName{Name: "PulseChat Support"}).To_Help_SupportName(), nil
}
