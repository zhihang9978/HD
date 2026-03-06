package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/webpage/internal/core"
)

// MessagesGetWebPagePreview570D6F6F
func (s *Service) MessagesGetWebPagePreview570D6F6F(ctx context.Context, request *mtproto.TLMessagesGetWebPagePreview570D6F6F) (*mtproto.Messages_WebPagePreview, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetWebPagePreview570D6F6F - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetWebPagePreview570D6F6F(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetWebPagePreview570D6F6F - reply: {%s}", r)
	return r, err
}
// MessagesGetWebPage8D9692A3
func (s *Service) MessagesGetWebPage8D9692A3(ctx context.Context, request *mtproto.TLMessagesGetWebPage8D9692A3) (*mtproto.Messages_WebPage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetWebPage8D9692A3 - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetWebPage8D9692A3(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetWebPage8D9692A3 - reply: {%s}", r)
	return r, err
}
// MessagesGetWebPagePreview8B68B0CC
func (s *Service) MessagesGetWebPagePreview8B68B0CC(ctx context.Context, request *mtproto.TLMessagesGetWebPagePreview8B68B0CC) (*mtproto.MessageMedia, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetWebPagePreview8B68B0CC - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetWebPagePreview8B68B0CC(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetWebPagePreview8B68B0CC - reply: {%s}", r)
	return r, err
}
// MessagesGetWebPage32CA8F91
func (s *Service) MessagesGetWebPage32CA8F91(ctx context.Context, request *mtproto.TLMessagesGetWebPage32CA8F91) (*mtproto.WebPage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetWebPage32CA8F91 - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetWebPage32CA8F91(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetWebPage32CA8F91 - reply: {%s}", r)
	return r, err
}
