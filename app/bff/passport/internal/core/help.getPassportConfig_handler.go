package core

import "github.com/teamgram/proto/mtproto"

func (c *PassportCore) HelpGetPassportConfig(in *mtproto.TLHelpGetPassportConfig) (*mtproto.Help_PassportConfig, error) {
	return mtproto.MakeTLHelpPassportConfigNotModified(&mtproto.Help_PassportConfig{}).To_Help_PassportConfig(), nil
}
