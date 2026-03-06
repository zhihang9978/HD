package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/reports/internal/core"
)

// AccountReportPeer
func (s *Service) AccountReportPeer(ctx context.Context, request *mtproto.TLAccountReportPeer) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountReportPeer - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountReportPeer(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountReportPeer - reply: {%s}", r)
	return r, err
}
// AccountReportProfilePhoto
func (s *Service) AccountReportProfilePhoto(ctx context.Context, request *mtproto.TLAccountReportProfilePhoto) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountReportProfilePhoto - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountReportProfilePhoto(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountReportProfilePhoto - reply: {%s}", r)
	return r, err
}
// MessagesReportSpam
func (s *Service) MessagesReportSpam(ctx context.Context, request *mtproto.TLMessagesReportSpam) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesReportSpam - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesReportSpam(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesReportSpam - reply: {%s}", r)
	return r, err
}
// MessagesReportFC78AF9B
func (s *Service) MessagesReportFC78AF9B(ctx context.Context, request *mtproto.TLMessagesReportFC78AF9B) (*mtproto.ReportResult, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesReportFC78AF9B - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesReportFC78AF9B(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesReportFC78AF9B - reply: {%s}", r)
	return r, err
}
// MessagesReportEncryptedSpam
func (s *Service) MessagesReportEncryptedSpam(ctx context.Context, request *mtproto.TLMessagesReportEncryptedSpam) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesReportEncryptedSpam - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesReportEncryptedSpam(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesReportEncryptedSpam - reply: {%s}", r)
	return r, err
}
// ChannelsReportSpam
func (s *Service) ChannelsReportSpam(ctx context.Context, request *mtproto.TLChannelsReportSpam) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("ChannelsReportSpam - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.ChannelsReportSpam(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("ChannelsReportSpam - reply: {%s}", r)
	return r, err
}
// MessagesReport8953AB4E
func (s *Service) MessagesReport8953AB4E(ctx context.Context, request *mtproto.TLMessagesReport8953AB4E) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesReport8953AB4E - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesReport8953AB4E(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesReport8953AB4E - reply: {%s}", r)
	return r, err
}
