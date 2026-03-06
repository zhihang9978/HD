package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/seamless/internal/core"
)

// AccountGetWebAuthorizations
func (s *Service) AccountGetWebAuthorizations(ctx context.Context, request *mtproto.TLAccountGetWebAuthorizations) (*mtproto.Account_WebAuthorizations, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountGetWebAuthorizations - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetWebAuthorizations(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountGetWebAuthorizations - reply: {%s}", r)
	return r, err
}
// AccountResetWebAuthorization
func (s *Service) AccountResetWebAuthorization(ctx context.Context, request *mtproto.TLAccountResetWebAuthorization) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountResetWebAuthorization - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountResetWebAuthorization(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountResetWebAuthorization - reply: {%s}", r)
	return r, err
}
// AccountResetWebAuthorizations
func (s *Service) AccountResetWebAuthorizations(ctx context.Context, request *mtproto.TLAccountResetWebAuthorizations) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountResetWebAuthorizations - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountResetWebAuthorizations(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountResetWebAuthorizations - reply: {%s}", r)
	return r, err
}
// MessagesRequestUrlAuth
func (s *Service) MessagesRequestUrlAuth(ctx context.Context, request *mtproto.TLMessagesRequestUrlAuth) (*mtproto.UrlAuthResult, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesRequestUrlAuth - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesRequestUrlAuth(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesRequestUrlAuth - reply: {%s}", r)
	return r, err
}
// MessagesAcceptUrlAuth
func (s *Service) MessagesAcceptUrlAuth(ctx context.Context, request *mtproto.TLMessagesAcceptUrlAuth) (*mtproto.UrlAuthResult, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesAcceptUrlAuth - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesAcceptUrlAuth(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesAcceptUrlAuth - reply: {%s}", r)
	return r, err
}
