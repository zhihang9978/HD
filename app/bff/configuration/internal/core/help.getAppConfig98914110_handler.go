package core

import "github.com/teamgram/proto/mtproto"

func (c *ConfigurationCore) HelpGetAppConfig98914110(in *mtproto.TLHelpGetAppConfig98914110) (*mtproto.JSONValue, error) {
	return mtproto.MakeTLJsonObject(&mtproto.JSONValue{Value_VECTORJSONOBJECTVALUE: []*mtproto.JSONObjectValue{}}).To_JSONValue(), nil
}
