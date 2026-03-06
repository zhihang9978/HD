package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/langpack/internal/core"
)

// LangpackGetLangPack
func (s *Service) LangpackGetLangPack(ctx context.Context, request *mtproto.TLLangpackGetLangPack) (*mtproto.LangPackDifference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("LangpackGetLangPack - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.LangpackGetLangPack(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("LangpackGetLangPack - reply: {%s}", r)
	return r, err
}
// LangpackGetStrings
func (s *Service) LangpackGetStrings(ctx context.Context, request *mtproto.TLLangpackGetStrings) (*mtproto.Vector_LangPackString, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("LangpackGetStrings - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.LangpackGetStrings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("LangpackGetStrings - reply: {%s}", r)
	return r, err
}
// LangpackGetDifference
func (s *Service) LangpackGetDifference(ctx context.Context, request *mtproto.TLLangpackGetDifference) (*mtproto.LangPackDifference, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("LangpackGetDifference - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.LangpackGetDifference(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("LangpackGetDifference - reply: {%s}", r)
	return r, err
}
// LangpackGetLanguages
func (s *Service) LangpackGetLanguages(ctx context.Context, request *mtproto.TLLangpackGetLanguages) (*mtproto.Vector_LangPackLanguage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("LangpackGetLanguages - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.LangpackGetLanguages(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("LangpackGetLanguages - reply: {%s}", r)
	return r, err
}
// LangpackGetLanguage
func (s *Service) LangpackGetLanguage(ctx context.Context, request *mtproto.TLLangpackGetLanguage) (*mtproto.LangPackLanguage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("LangpackGetLanguage - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.LangpackGetLanguage(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("LangpackGetLanguage - reply: {%s}", r)
	return r, err
}
