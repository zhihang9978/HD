package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/themes/internal/core"
)

// AccountUploadTheme
func (s *Service) AccountUploadTheme(ctx context.Context, request *mtproto.TLAccountUploadTheme) (*mtproto.Document, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountUploadTheme - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountUploadTheme(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountUploadTheme - reply: {%s}", r)
	return r, err
}
// AccountCreateTheme
func (s *Service) AccountCreateTheme(ctx context.Context, request *mtproto.TLAccountCreateTheme) (*mtproto.Theme, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountCreateTheme - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountCreateTheme(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountCreateTheme - reply: {%s}", r)
	return r, err
}
// AccountUpdateTheme
func (s *Service) AccountUpdateTheme(ctx context.Context, request *mtproto.TLAccountUpdateTheme) (*mtproto.Theme, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountUpdateTheme - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountUpdateTheme(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountUpdateTheme - reply: {%s}", r)
	return r, err
}
// AccountSaveTheme
func (s *Service) AccountSaveTheme(ctx context.Context, request *mtproto.TLAccountSaveTheme) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountSaveTheme - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountSaveTheme(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountSaveTheme - reply: {%s}", r)
	return r, err
}
// AccountInstallTheme
func (s *Service) AccountInstallTheme(ctx context.Context, request *mtproto.TLAccountInstallTheme) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountInstallTheme - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountInstallTheme(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountInstallTheme - reply: {%s}", r)
	return r, err
}
// AccountGetTheme
func (s *Service) AccountGetTheme(ctx context.Context, request *mtproto.TLAccountGetTheme) (*mtproto.Theme, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountGetTheme - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetTheme(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountGetTheme - reply: {%s}", r)
	return r, err
}
// AccountGetThemes
func (s *Service) AccountGetThemes(ctx context.Context, request *mtproto.TLAccountGetThemes) (*mtproto.Account_Themes, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountGetThemes - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetThemes(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountGetThemes - reply: {%s}", r)
	return r, err
}
// AccountGetChatThemes
func (s *Service) AccountGetChatThemes(ctx context.Context, request *mtproto.TLAccountGetChatThemes) (*mtproto.Account_Themes, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountGetChatThemes - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetChatThemes(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountGetChatThemes - reply: {%s}", r)
	return r, err
}
// AccountGetUniqueGiftChatThemes
func (s *Service) AccountGetUniqueGiftChatThemes(ctx context.Context, request *mtproto.TLAccountGetUniqueGiftChatThemes) (*mtproto.Account_ChatThemes, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountGetUniqueGiftChatThemes - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetUniqueGiftChatThemes(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountGetUniqueGiftChatThemes - reply: {%s}", r)
	return r, err
}
// MessagesSetChatTheme
func (s *Service) MessagesSetChatTheme(ctx context.Context, request *mtproto.TLMessagesSetChatTheme) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSetChatTheme - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSetChatTheme(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSetChatTheme - reply: {%s}", r)
	return r, err
}
