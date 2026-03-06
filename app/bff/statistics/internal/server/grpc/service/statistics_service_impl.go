package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/statistics/internal/core"
)

// StatsGetBroadcastStats
func (s *Service) StatsGetBroadcastStats(ctx context.Context, request *mtproto.TLStatsGetBroadcastStats) (*mtproto.Stats_BroadcastStats, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StatsGetBroadcastStats - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StatsGetBroadcastStats(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StatsGetBroadcastStats - reply: {%s}", r)
	return r, err
}
// StatsLoadAsyncGraph
func (s *Service) StatsLoadAsyncGraph(ctx context.Context, request *mtproto.TLStatsLoadAsyncGraph) (*mtproto.StatsGraph, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StatsLoadAsyncGraph - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StatsLoadAsyncGraph(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StatsLoadAsyncGraph - reply: {%s}", r)
	return r, err
}
// StatsGetMegagroupStats
func (s *Service) StatsGetMegagroupStats(ctx context.Context, request *mtproto.TLStatsGetMegagroupStats) (*mtproto.Stats_MegagroupStats, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StatsGetMegagroupStats - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StatsGetMegagroupStats(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StatsGetMegagroupStats - reply: {%s}", r)
	return r, err
}
// StatsGetMessagePublicForwards5F150144
func (s *Service) StatsGetMessagePublicForwards5F150144(ctx context.Context, request *mtproto.TLStatsGetMessagePublicForwards5F150144) (*mtproto.Stats_PublicForwards, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StatsGetMessagePublicForwards5F150144 - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StatsGetMessagePublicForwards5F150144(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StatsGetMessagePublicForwards5F150144 - reply: {%s}", r)
	return r, err
}
// StatsGetMessageStats
func (s *Service) StatsGetMessageStats(ctx context.Context, request *mtproto.TLStatsGetMessageStats) (*mtproto.Stats_MessageStats, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StatsGetMessageStats - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StatsGetMessageStats(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StatsGetMessageStats - reply: {%s}", r)
	return r, err
}
// StatsGetStoryStats
func (s *Service) StatsGetStoryStats(ctx context.Context, request *mtproto.TLStatsGetStoryStats) (*mtproto.Stats_StoryStats, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StatsGetStoryStats - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StatsGetStoryStats(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StatsGetStoryStats - reply: {%s}", r)
	return r, err
}
// StatsGetStoryPublicForwards
func (s *Service) StatsGetStoryPublicForwards(ctx context.Context, request *mtproto.TLStatsGetStoryPublicForwards) (*mtproto.Stats_PublicForwards, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StatsGetStoryPublicForwards - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StatsGetStoryPublicForwards(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StatsGetStoryPublicForwards - reply: {%s}", r)
	return r, err
}
// StatsGetMessagePublicForwards5630281B
func (s *Service) StatsGetMessagePublicForwards5630281B(ctx context.Context, request *mtproto.TLStatsGetMessagePublicForwards5630281B) (*mtproto.Messages_Messages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("StatsGetMessagePublicForwards5630281B - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.StatsGetMessagePublicForwards5630281B(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("StatsGetMessagePublicForwards5630281B - reply: {%s}", r)
	return r, err
}
