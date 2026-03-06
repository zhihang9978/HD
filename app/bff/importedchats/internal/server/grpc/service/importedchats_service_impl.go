package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/importedchats/internal/core"
)

// MessagesCheckHistoryImport
func (s *Service) MessagesCheckHistoryImport(ctx context.Context, request *mtproto.TLMessagesCheckHistoryImport) (*mtproto.Messages_HistoryImportParsed, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesCheckHistoryImport - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesCheckHistoryImport(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesCheckHistoryImport - reply: {%s}", r)
	return r, err
}
// MessagesInitHistoryImport
func (s *Service) MessagesInitHistoryImport(ctx context.Context, request *mtproto.TLMessagesInitHistoryImport) (*mtproto.Messages_HistoryImport, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesInitHistoryImport - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesInitHistoryImport(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesInitHistoryImport - reply: {%s}", r)
	return r, err
}
// MessagesUploadImportedMedia
func (s *Service) MessagesUploadImportedMedia(ctx context.Context, request *mtproto.TLMessagesUploadImportedMedia) (*mtproto.MessageMedia, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesUploadImportedMedia - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesUploadImportedMedia(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesUploadImportedMedia - reply: {%s}", r)
	return r, err
}
// MessagesStartHistoryImport
func (s *Service) MessagesStartHistoryImport(ctx context.Context, request *mtproto.TLMessagesStartHistoryImport) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesStartHistoryImport - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesStartHistoryImport(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesStartHistoryImport - reply: {%s}", r)
	return r, err
}
// MessagesCheckHistoryImportPeer
func (s *Service) MessagesCheckHistoryImportPeer(ctx context.Context, request *mtproto.TLMessagesCheckHistoryImportPeer) (*mtproto.Messages_CheckedHistoryImportPeer, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesCheckHistoryImportPeer - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesCheckHistoryImportPeer(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesCheckHistoryImportPeer - reply: {%s}", r)
	return r, err
}
