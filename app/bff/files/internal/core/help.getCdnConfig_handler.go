package core

import "github.com/teamgram/proto/mtproto"

func (c *FilesCore) HelpGetCdnConfig(in *mtproto.TLHelpGetCdnConfig) (*mtproto.CdnConfig, error) {
	return mtproto.MakeTLCdnConfig(&mtproto.CdnConfig{PublicKeys: []*mtproto.CdnPublicKey{}}).To_CdnConfig(), nil
}
