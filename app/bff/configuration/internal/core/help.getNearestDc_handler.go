package core

import "github.com/teamgram/proto/mtproto"

func (c *ConfigurationCore) HelpGetNearestDc(in *mtproto.TLHelpGetNearestDc) (*mtproto.NearestDc, error) {
	return mtproto.MakeTLNearestDc(&mtproto.NearestDc{Country: "CN", ThisDc: 1, NearestDc: 1}).To_NearestDc(), nil
}
