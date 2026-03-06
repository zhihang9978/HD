package core

import (
	"github.com/teamgram/proto/mtproto"
)

// HelpGetPromoData
func (c *PromoDataCore) HelpGetPromoData(in *mtproto.TLHelpGetPromoData) (*mtproto.Help_PromoData, error) {
	c.Logger.Errorf("HelpGetPromoData - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
