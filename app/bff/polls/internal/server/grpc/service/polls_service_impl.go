package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/polls/internal/core"
)

// MessagesSendVote
func (s *Service) MessagesSendVote(ctx context.Context, request *mtproto.TLMessagesSendVote) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSendVote - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSendVote(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSendVote - reply: {%s}", r)
	return r, err
}
// MessagesGetPollResults
func (s *Service) MessagesGetPollResults(ctx context.Context, request *mtproto.TLMessagesGetPollResults) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetPollResults - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetPollResults(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetPollResults - reply: {%s}", r)
	return r, err
}
// MessagesGetPollVotes
func (s *Service) MessagesGetPollVotes(ctx context.Context, request *mtproto.TLMessagesGetPollVotes) (*mtproto.Messages_VotesList, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetPollVotes - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetPollVotes(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetPollVotes - reply: {%s}", r)
	return r, err
}
