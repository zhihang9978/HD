package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/deeplinks/internal/core"
)

// MessagesStartBot
func (s *Service) MessagesStartBot(ctx context.Context, request *mtproto.TLMessagesStartBot) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesStartBot - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesStartBot(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesStartBot - reply: {%s}", r)
	return r, err
}
// HelpGetRecentMeUrls
func (s *Service) HelpGetRecentMeUrls(ctx context.Context, request *mtproto.TLHelpGetRecentMeUrls) (*mtproto.Help_RecentMeUrls, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("HelpGetRecentMeUrls - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpGetRecentMeUrls(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("HelpGetRecentMeUrls - reply: {%s}", r)
	return r, err
}
// HelpGetDeepLinkInfo
func (s *Service) HelpGetDeepLinkInfo(ctx context.Context, request *mtproto.TLHelpGetDeepLinkInfo) (*mtproto.Help_DeepLinkInfo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("HelpGetDeepLinkInfo - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpGetDeepLinkInfo(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("HelpGetDeepLinkInfo - reply: {%s}", r)
	return r, err
}
