package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/gifs/internal/core"
)

// MessagesGetSavedGifs
func (s *Service) MessagesGetSavedGifs(ctx context.Context, request *mtproto.TLMessagesGetSavedGifs) (*mtproto.Messages_SavedGifs, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetSavedGifs - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetSavedGifs(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetSavedGifs - reply: {%s}", r)
	return r, err
}
// MessagesSaveGif
func (s *Service) MessagesSaveGif(ctx context.Context, request *mtproto.TLMessagesSaveGif) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSaveGif - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSaveGif(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSaveGif - reply: {%s}", r)
	return r, err
}
