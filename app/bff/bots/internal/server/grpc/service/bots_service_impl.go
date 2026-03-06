package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/bots/internal/core"
)

// BotsSetBotCommands
func (s *Service) BotsSetBotCommands(ctx context.Context, request *mtproto.TLBotsSetBotCommands) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("BotsSetBotCommands - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.BotsSetBotCommands(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("BotsSetBotCommands - reply: {%s}", r)
	return r, err
}
// BotsResetBotCommands
func (s *Service) BotsResetBotCommands(ctx context.Context, request *mtproto.TLBotsResetBotCommands) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("BotsResetBotCommands - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.BotsResetBotCommands(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("BotsResetBotCommands - reply: {%s}", r)
	return r, err
}
// BotsGetBotCommands
func (s *Service) BotsGetBotCommands(ctx context.Context, request *mtproto.TLBotsGetBotCommands) (*mtproto.Vector_BotCommand, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("BotsGetBotCommands - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.BotsGetBotCommands(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("BotsGetBotCommands - reply: {%s}", r)
	return r, err
}
// BotsSetBotInfo
func (s *Service) BotsSetBotInfo(ctx context.Context, request *mtproto.TLBotsSetBotInfo) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("BotsSetBotInfo - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.BotsSetBotInfo(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("BotsSetBotInfo - reply: {%s}", r)
	return r, err
}
// BotsGetBotInfoDCD914FD
func (s *Service) BotsGetBotInfoDCD914FD(ctx context.Context, request *mtproto.TLBotsGetBotInfoDCD914FD) (*mtproto.Bots_BotInfo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("BotsGetBotInfoDCD914FD - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.BotsGetBotInfoDCD914FD(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("BotsGetBotInfoDCD914FD - reply: {%s}", r)
	return r, err
}
// BotsGetAdminedBots
func (s *Service) BotsGetAdminedBots(ctx context.Context, request *mtproto.TLBotsGetAdminedBots) (*mtproto.Vector_User, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("BotsGetAdminedBots - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.BotsGetAdminedBots(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("BotsGetAdminedBots - reply: {%s}", r)
	return r, err
}
// BotsGetBotInfo75EC12E6
func (s *Service) BotsGetBotInfo75EC12E6(ctx context.Context, request *mtproto.TLBotsGetBotInfo75EC12E6) (*mtproto.Vector_String, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("BotsGetBotInfo75EC12E6 - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.BotsGetBotInfo75EC12E6(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("BotsGetBotInfo75EC12E6 - reply: {%s}", r)
	return r, err
}
