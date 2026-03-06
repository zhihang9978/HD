package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/emoji/internal/core"
)

// MessagesGetEmojiKeywords
func (s *Service) MessagesGetEmojiKeywords(ctx context.Context, request *mtproto.TLMessagesGetEmojiKeywords) (*mtproto.EmojiKeywordsDifference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetEmojiKeywords - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetEmojiKeywords(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetEmojiKeywords - reply: {%s}", r)
	return r, err
}
// MessagesGetEmojiKeywordsDifference
func (s *Service) MessagesGetEmojiKeywordsDifference(ctx context.Context, request *mtproto.TLMessagesGetEmojiKeywordsDifference) (*mtproto.EmojiKeywordsDifference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetEmojiKeywordsDifference - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetEmojiKeywordsDifference(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetEmojiKeywordsDifference - reply: {%s}", r)
	return r, err
}
// MessagesGetEmojiKeywordsLanguages
func (s *Service) MessagesGetEmojiKeywordsLanguages(ctx context.Context, request *mtproto.TLMessagesGetEmojiKeywordsLanguages) (*mtproto.Vector_EmojiLanguage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetEmojiKeywordsLanguages - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetEmojiKeywordsLanguages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetEmojiKeywordsLanguages - reply: {%s}", r)
	return r, err
}
// MessagesGetEmojiURL
func (s *Service) MessagesGetEmojiURL(ctx context.Context, request *mtproto.TLMessagesGetEmojiURL) (*mtproto.EmojiURL, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetEmojiURL - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetEmojiURL(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetEmojiURL - reply: {%s}", r)
	return r, err
}
