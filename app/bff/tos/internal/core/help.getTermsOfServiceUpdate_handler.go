package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

func (c *TosCore) HelpGetTermsOfServiceUpdate(in *mtproto.TLHelpGetTermsOfServiceUpdate) (*mtproto.Help_TermsOfServiceUpdate, error) {
	return mtproto.MakeTLHelpTermsOfServiceUpdateEmpty(&mtproto.Help_TermsOfServiceUpdate{
		Expires: int32(time.Now().Unix() + 86400),
	}).To_Help_TermsOfServiceUpdate(), nil
}
