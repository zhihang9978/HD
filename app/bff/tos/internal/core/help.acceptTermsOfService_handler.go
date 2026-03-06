package core

import "github.com/teamgram/proto/mtproto"

func (c *TosCore) HelpAcceptTermsOfService(in *mtproto.TLHelpAcceptTermsOfService) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
