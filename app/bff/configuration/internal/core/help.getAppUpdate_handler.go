package core

import "github.com/teamgram/proto/mtproto"

func (c *ConfigurationCore) HelpGetAppUpdate(in *mtproto.TLHelpGetAppUpdate) (*mtproto.Help_AppUpdate, error) {
	return mtproto.MakeTLHelpNoAppUpdate(&mtproto.Help_AppUpdate{}).To_Help_AppUpdate(), nil
}
