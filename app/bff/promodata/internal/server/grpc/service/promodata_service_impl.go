package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/promodata/internal/core"
)

// HelpGetPromoData
func (s *Service) HelpGetPromoData(ctx context.Context, request *mtproto.TLHelpGetPromoData) (*mtproto.Help_PromoData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("HelpGetPromoData - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpGetPromoData(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("HelpGetPromoData - reply: {%s}", r)
	return r, err
}
// HelpHidePromoData
func (s *Service) HelpHidePromoData(ctx context.Context, request *mtproto.TLHelpHidePromoData) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("HelpHidePromoData - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpHidePromoData(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("HelpHidePromoData - reply: {%s}", r)
	return r, err
}
