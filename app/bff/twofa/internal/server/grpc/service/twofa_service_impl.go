package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/twofa/internal/core"
)

// AccountGetPassword
func (s *Service) AccountGetPassword(ctx context.Context, request *mtproto.TLAccountGetPassword) (*mtproto.Account_Password, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountGetPassword - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetPassword(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountGetPassword - reply: {%s}", r)
	return r, err
}
// AccountGetPasswordSettings
func (s *Service) AccountGetPasswordSettings(ctx context.Context, request *mtproto.TLAccountGetPasswordSettings) (*mtproto.Account_PasswordSettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountGetPasswordSettings - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetPasswordSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountGetPasswordSettings - reply: {%s}", r)
	return r, err
}
// AccountUpdatePasswordSettings
func (s *Service) AccountUpdatePasswordSettings(ctx context.Context, request *mtproto.TLAccountUpdatePasswordSettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountUpdatePasswordSettings - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountUpdatePasswordSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountUpdatePasswordSettings - reply: {%s}", r)
	return r, err
}
// AccountConfirmPasswordEmail
func (s *Service) AccountConfirmPasswordEmail(ctx context.Context, request *mtproto.TLAccountConfirmPasswordEmail) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountConfirmPasswordEmail - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountConfirmPasswordEmail(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountConfirmPasswordEmail - reply: {%s}", r)
	return r, err
}
// AccountResendPasswordEmail
func (s *Service) AccountResendPasswordEmail(ctx context.Context, request *mtproto.TLAccountResendPasswordEmail) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountResendPasswordEmail - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountResendPasswordEmail(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountResendPasswordEmail - reply: {%s}", r)
	return r, err
}
// AccountCancelPasswordEmail
func (s *Service) AccountCancelPasswordEmail(ctx context.Context, request *mtproto.TLAccountCancelPasswordEmail) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountCancelPasswordEmail - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountCancelPasswordEmail(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountCancelPasswordEmail - reply: {%s}", r)
	return r, err
}
// AccountDeclinePasswordReset
func (s *Service) AccountDeclinePasswordReset(ctx context.Context, request *mtproto.TLAccountDeclinePasswordReset) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountDeclinePasswordReset - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountDeclinePasswordReset(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountDeclinePasswordReset - reply: {%s}", r)
	return r, err
}
