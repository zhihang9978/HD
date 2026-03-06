package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/reactions/internal/core"
)

// MessagesSendReaction
func (s *Service) MessagesSendReaction(ctx context.Context, request *mtproto.TLMessagesSendReaction) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSendReaction - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSendReaction(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSendReaction - reply: {%s}", r)
	return r, err
}
// MessagesGetMessagesReactions
func (s *Service) MessagesGetMessagesReactions(ctx context.Context, request *mtproto.TLMessagesGetMessagesReactions) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetMessagesReactions - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetMessagesReactions(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetMessagesReactions - reply: {%s}", r)
	return r, err
}
// MessagesGetMessageReactionsList
func (s *Service) MessagesGetMessageReactionsList(ctx context.Context, request *mtproto.TLMessagesGetMessageReactionsList) (*mtproto.Messages_MessageReactionsList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetMessageReactionsList - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetMessageReactionsList(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetMessageReactionsList - reply: {%s}", r)
	return r, err
}
// MessagesSetChatAvailableReactions
func (s *Service) MessagesSetChatAvailableReactions(ctx context.Context, request *mtproto.TLMessagesSetChatAvailableReactions) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSetChatAvailableReactions - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSetChatAvailableReactions(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSetChatAvailableReactions - reply: {%s}", r)
	return r, err
}
// MessagesGetAvailableReactions
func (s *Service) MessagesGetAvailableReactions(ctx context.Context, request *mtproto.TLMessagesGetAvailableReactions) (*mtproto.Messages_AvailableReactions, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetAvailableReactions - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetAvailableReactions(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetAvailableReactions - reply: {%s}", r)
	return r, err
}
// MessagesSetDefaultReaction
func (s *Service) MessagesSetDefaultReaction(ctx context.Context, request *mtproto.TLMessagesSetDefaultReaction) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSetDefaultReaction - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSetDefaultReaction(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSetDefaultReaction - reply: {%s}", r)
	return r, err
}
// MessagesGetUnreadReactions
func (s *Service) MessagesGetUnreadReactions(ctx context.Context, request *mtproto.TLMessagesGetUnreadReactions) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetUnreadReactions - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetUnreadReactions(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetUnreadReactions - reply: {%s}", r)
	return r, err
}
// MessagesReadReactions
func (s *Service) MessagesReadReactions(ctx context.Context, request *mtproto.TLMessagesReadReactions) (*mtproto.Messages_AffectedHistory, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesReadReactions - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesReadReactions(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesReadReactions - reply: {%s}", r)
	return r, err
}
// MessagesReportReaction
func (s *Service) MessagesReportReaction(ctx context.Context, request *mtproto.TLMessagesReportReaction) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesReportReaction - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesReportReaction(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesReportReaction - reply: {%s}", r)
	return r, err
}
// MessagesGetTopReactions
func (s *Service) MessagesGetTopReactions(ctx context.Context, request *mtproto.TLMessagesGetTopReactions) (*mtproto.Messages_Reactions, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetTopReactions - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetTopReactions(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetTopReactions - reply: {%s}", r)
	return r, err
}
// MessagesGetRecentReactions
func (s *Service) MessagesGetRecentReactions(ctx context.Context, request *mtproto.TLMessagesGetRecentReactions) (*mtproto.Messages_Reactions, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetRecentReactions - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetRecentReactions(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetRecentReactions - reply: {%s}", r)
	return r, err
}
// MessagesClearRecentReactions
func (s *Service) MessagesClearRecentReactions(ctx context.Context, request *mtproto.TLMessagesClearRecentReactions) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesClearRecentReactions - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesClearRecentReactions(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesClearRecentReactions - reply: {%s}", r)
	return r, err
}
// MessagesSendPaidReaction
func (s *Service) MessagesSendPaidReaction(ctx context.Context, request *mtproto.TLMessagesSendPaidReaction) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSendPaidReaction - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSendPaidReaction(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSendPaidReaction - reply: {%s}", r)
	return r, err
}
// MessagesTogglePaidReactionPrivacy
func (s *Service) MessagesTogglePaidReactionPrivacy(ctx context.Context, request *mtproto.TLMessagesTogglePaidReactionPrivacy) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesTogglePaidReactionPrivacy - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesTogglePaidReactionPrivacy(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesTogglePaidReactionPrivacy - reply: {%s}", r)
	return r, err
}
// MessagesGetPaidReactionPrivacy
func (s *Service) MessagesGetPaidReactionPrivacy(ctx context.Context, request *mtproto.TLMessagesGetPaidReactionPrivacy) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetPaidReactionPrivacy - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetPaidReactionPrivacy(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetPaidReactionPrivacy - reply: {%s}", r)
	return r, err
}
