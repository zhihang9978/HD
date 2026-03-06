package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/scheduledmessages/internal/core"
)

// MessagesGetScheduledHistory
func (s *Service) MessagesGetScheduledHistory(ctx context.Context, request *mtproto.TLMessagesGetScheduledHistory) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetScheduledHistory - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetScheduledHistory(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetScheduledHistory - reply: {%s}", r)
	return r, err
}
// MessagesGetScheduledMessages
func (s *Service) MessagesGetScheduledMessages(ctx context.Context, request *mtproto.TLMessagesGetScheduledMessages) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetScheduledMessages - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetScheduledMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetScheduledMessages - reply: {%s}", r)
	return r, err
}
// MessagesSendScheduledMessages
func (s *Service) MessagesSendScheduledMessages(ctx context.Context, request *mtproto.TLMessagesSendScheduledMessages) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSendScheduledMessages - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSendScheduledMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSendScheduledMessages - reply: {%s}", r)
	return r, err
}
// MessagesDeleteScheduledMessages
func (s *Service) MessagesDeleteScheduledMessages(ctx context.Context, request *mtproto.TLMessagesDeleteScheduledMessages) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesDeleteScheduledMessages - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesDeleteScheduledMessages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesDeleteScheduledMessages - reply: {%s}", r)
	return r, err
}
