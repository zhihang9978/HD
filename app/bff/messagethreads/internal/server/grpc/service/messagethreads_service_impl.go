package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/messagethreads/internal/core"
)

// ContactsBlockFromReplies
func (s *Service) ContactsBlockFromReplies(ctx context.Context, request *mtproto.TLContactsBlockFromReplies) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("ContactsBlockFromReplies - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ContactsBlockFromReplies(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("ContactsBlockFromReplies - reply: {%s}", r)
	return r, err
}
// MessagesGetReplies
func (s *Service) MessagesGetReplies(ctx context.Context, request *mtproto.TLMessagesGetReplies) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetReplies - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetReplies(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetReplies - reply: {%s}", r)
	return r, err
}
// MessagesGetDiscussionMessage
func (s *Service) MessagesGetDiscussionMessage(ctx context.Context, request *mtproto.TLMessagesGetDiscussionMessage) (*mtproto.Messages_DiscussionMessage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetDiscussionMessage - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetDiscussionMessage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetDiscussionMessage - reply: {%s}", r)
	return r, err
}
// MessagesReadDiscussion
func (s *Service) MessagesReadDiscussion(ctx context.Context, request *mtproto.TLMessagesReadDiscussion) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesReadDiscussion - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesReadDiscussion(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesReadDiscussion - reply: {%s}", r)
	return r, err
}
