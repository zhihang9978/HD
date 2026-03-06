package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/internalbot/internal/core"
)

// HelpSetBotUpdatesStatus
func (s *Service) HelpSetBotUpdatesStatus(ctx context.Context, request *mtproto.TLHelpSetBotUpdatesStatus) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("HelpSetBotUpdatesStatus - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpSetBotUpdatesStatus(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("HelpSetBotUpdatesStatus - reply: {%s}", r)
	return r, err
}
// BotsSendCustomRequest
func (s *Service) BotsSendCustomRequest(ctx context.Context, request *mtproto.TLBotsSendCustomRequest) (*mtproto.DataJSON, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("BotsSendCustomRequest - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.BotsSendCustomRequest(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("BotsSendCustomRequest - reply: {%s}", r)
	return r, err
}
// BotsAnswerWebhookJSONQuery
func (s *Service) BotsAnswerWebhookJSONQuery(ctx context.Context, request *mtproto.TLBotsAnswerWebhookJSONQuery) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("BotsAnswerWebhookJSONQuery - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.BotsAnswerWebhookJSONQuery(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("BotsAnswerWebhookJSONQuery - reply: {%s}", r)
	return r, err
}
