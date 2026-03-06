package core

import (
	"github.com/teamgram/proto/mtproto"
)

// HelpHidePromoData
func (c *PromoDataCore) HelpHidePromoData(in *mtproto.TLHelpHidePromoData) (*mtproto.Bool, error) {
	c.Logger.Errorf("HelpHidePromoData - not impl")

	return mtproto.BoolTrue, nil
}
