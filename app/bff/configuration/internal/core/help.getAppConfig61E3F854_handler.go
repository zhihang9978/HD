package core

import "github.com/teamgram/proto/mtproto"

func (c *ConfigurationCore) HelpGetAppConfig61E3F854(in *mtproto.TLHelpGetAppConfig61E3F854) (*mtproto.Help_AppConfig, error) {
	return mtproto.MakeTLHelpAppConfig(&mtproto.Help_AppConfig{
		Hash: 0,
		Config: mtproto.MakeTLJsonObject(&mtproto.JSONValue{Value_VECTORJSONOBJECTVALUE: []*mtproto.JSONObjectValue{}}).To_JSONValue(),
	}).To_Help_AppConfig(), nil
}
