package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/games/internal/core"
)

// MessagesSetGameScore
func (s *Service) MessagesSetGameScore(ctx context.Context, request *mtproto.TLMessagesSetGameScore) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSetGameScore - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSetGameScore(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSetGameScore - reply: {%s}", r)
	return r, err
}
// MessagesSetInlineGameScore
func (s *Service) MessagesSetInlineGameScore(ctx context.Context, request *mtproto.TLMessagesSetInlineGameScore) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSetInlineGameScore - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSetInlineGameScore(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSetInlineGameScore - reply: {%s}", r)
	return r, err
}
// MessagesGetGameHighScores
func (s *Service) MessagesGetGameHighScores(ctx context.Context, request *mtproto.TLMessagesGetGameHighScores) (*mtproto.Messages_HighScores, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetGameHighScores - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetGameHighScores(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetGameHighScores - reply: {%s}", r)
	return r, err
}
// MessagesGetInlineGameHighScores
func (s *Service) MessagesGetInlineGameHighScores(ctx context.Context, request *mtproto.TLMessagesGetInlineGameHighScores) (*mtproto.Messages_HighScores, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetInlineGameHighScores - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetInlineGameHighScores(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetInlineGameHighScores - reply: {%s}", r)
	return r, err
}
// MessagesGetEmojiGameInfo
func (s *Service) MessagesGetEmojiGameInfo(ctx context.Context, request *mtproto.TLMessagesGetEmojiGameInfo) (*mtproto.Messages_EmojiGameInfo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetEmojiGameInfo - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetEmojiGameInfo(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetEmojiGameInfo - reply: {%s}", r)
	return r, err
}
