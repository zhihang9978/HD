package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/wallpapers/internal/core"
)

// AccountGetWallPapers
func (s *Service) AccountGetWallPapers(ctx context.Context, request *mtproto.TLAccountGetWallPapers) (*mtproto.Account_WallPapers, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountGetWallPapers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetWallPapers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountGetWallPapers - reply: {%s}", r)
	return r, err
}
// AccountGetWallPaper
func (s *Service) AccountGetWallPaper(ctx context.Context, request *mtproto.TLAccountGetWallPaper) (*mtproto.WallPaper, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountGetWallPaper - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetWallPaper(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountGetWallPaper - reply: {%s}", r)
	return r, err
}
// AccountUploadWallPaper
func (s *Service) AccountUploadWallPaper(ctx context.Context, request *mtproto.TLAccountUploadWallPaper) (*mtproto.WallPaper, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountUploadWallPaper - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountUploadWallPaper(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountUploadWallPaper - reply: {%s}", r)
	return r, err
}
// AccountSaveWallPaper
func (s *Service) AccountSaveWallPaper(ctx context.Context, request *mtproto.TLAccountSaveWallPaper) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountSaveWallPaper - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountSaveWallPaper(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountSaveWallPaper - reply: {%s}", r)
	return r, err
}
// AccountInstallWallPaper
func (s *Service) AccountInstallWallPaper(ctx context.Context, request *mtproto.TLAccountInstallWallPaper) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountInstallWallPaper - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountInstallWallPaper(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountInstallWallPaper - reply: {%s}", r)
	return r, err
}
// AccountResetWallPapers
func (s *Service) AccountResetWallPapers(ctx context.Context, request *mtproto.TLAccountResetWallPapers) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountResetWallPapers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountResetWallPapers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountResetWallPapers - reply: {%s}", r)
	return r, err
}
// AccountGetMultiWallPapers
func (s *Service) AccountGetMultiWallPapers(ctx context.Context, request *mtproto.TLAccountGetMultiWallPapers) (*mtproto.Vector_WallPaper, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountGetMultiWallPapers - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetMultiWallPapers(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountGetMultiWallPapers - reply: {%s}", r)
	return r, err
}
// MessagesSetChatWallPaper
func (s *Service) MessagesSetChatWallPaper(ctx context.Context, request *mtproto.TLMessagesSetChatWallPaper) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSetChatWallPaper - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSetChatWallPaper(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSetChatWallPaper - reply: {%s}", r)
	return r, err
}
