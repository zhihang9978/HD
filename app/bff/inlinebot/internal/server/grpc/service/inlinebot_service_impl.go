package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/inlinebot/internal/core"
)

// MessagesGetInlineBotResults
func (s *Service) MessagesGetInlineBotResults(ctx context.Context, request *mtproto.TLMessagesGetInlineBotResults) (*mtproto.Messages_BotResults, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetInlineBotResults - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetInlineBotResults(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetInlineBotResults - reply: {%s}", r)
	return r, err
}
// MessagesSetInlineBotResults
func (s *Service) MessagesSetInlineBotResults(ctx context.Context, request *mtproto.TLMessagesSetInlineBotResults) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSetInlineBotResults - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSetInlineBotResults(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSetInlineBotResults - reply: {%s}", r)
	return r, err
}
// MessagesSendInlineBotResult
func (s *Service) MessagesSendInlineBotResult(ctx context.Context, request *mtproto.TLMessagesSendInlineBotResult) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSendInlineBotResult - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSendInlineBotResult(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSendInlineBotResult - reply: {%s}", r)
	return r, err
}
// MessagesEditInlineBotMessage
func (s *Service) MessagesEditInlineBotMessage(ctx context.Context, request *mtproto.TLMessagesEditInlineBotMessage) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesEditInlineBotMessage - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesEditInlineBotMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesEditInlineBotMessage - reply: {%s}", r)
	return r, err
}
// MessagesGetBotCallbackAnswer
func (s *Service) MessagesGetBotCallbackAnswer(ctx context.Context, request *mtproto.TLMessagesGetBotCallbackAnswer) (*mtproto.Messages_BotCallbackAnswer, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetBotCallbackAnswer - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetBotCallbackAnswer(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetBotCallbackAnswer - reply: {%s}", r)
	return r, err
}
// MessagesSetBotCallbackAnswer
func (s *Service) MessagesSetBotCallbackAnswer(ctx context.Context, request *mtproto.TLMessagesSetBotCallbackAnswer) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSetBotCallbackAnswer - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSetBotCallbackAnswer(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSetBotCallbackAnswer - reply: {%s}", r)
	return r, err
}
// MessagesSendBotRequestedPeer
func (s *Service) MessagesSendBotRequestedPeer(ctx context.Context, request *mtproto.TLMessagesSendBotRequestedPeer) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSendBotRequestedPeer - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSendBotRequestedPeer(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSendBotRequestedPeer - reply: {%s}", r)
	return r, err
}
