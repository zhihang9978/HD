package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/tsf/internal/core"
)

// HelpGetUserInfo
func (s *Service) HelpGetUserInfo(ctx context.Context, request *mtproto.TLHelpGetUserInfo) (*mtproto.Help_UserInfo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("HelpGetUserInfo - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpGetUserInfo(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("HelpGetUserInfo - reply: {%s}", r)
	return r, err
}
// HelpEditUserInfo
func (s *Service) HelpEditUserInfo(ctx context.Context, request *mtproto.TLHelpEditUserInfo) (*mtproto.Help_UserInfo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("HelpEditUserInfo - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.HelpEditUserInfo(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("HelpEditUserInfo - reply: {%s}", r)
	return r, err
}
