package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/voipcalls/internal/core"
)

// MessagesDeletePhoneCallHistory
func (s *Service) MessagesDeletePhoneCallHistory(ctx context.Context, request *mtproto.TLMessagesDeletePhoneCallHistory) (*mtproto.Messages_AffectedFoundMessages, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesDeletePhoneCallHistory - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesDeletePhoneCallHistory(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesDeletePhoneCallHistory - reply: {%s}", r)
	return r, err
}
// PhoneGetCallConfig
func (s *Service) PhoneGetCallConfig(ctx context.Context, request *mtproto.TLPhoneGetCallConfig) (*mtproto.DataJSON, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PhoneGetCallConfig - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PhoneGetCallConfig(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PhoneGetCallConfig - reply: {%s}", r)
	return r, err
}
// PhoneRequestCall
func (s *Service) PhoneRequestCall(ctx context.Context, request *mtproto.TLPhoneRequestCall) (*mtproto.Phone_PhoneCall, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PhoneRequestCall - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PhoneRequestCall(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PhoneRequestCall - reply: {%s}", r)
	return r, err
}
// PhoneAcceptCall
func (s *Service) PhoneAcceptCall(ctx context.Context, request *mtproto.TLPhoneAcceptCall) (*mtproto.Phone_PhoneCall, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PhoneAcceptCall - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PhoneAcceptCall(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PhoneAcceptCall - reply: {%s}", r)
	return r, err
}
// PhoneConfirmCall
func (s *Service) PhoneConfirmCall(ctx context.Context, request *mtproto.TLPhoneConfirmCall) (*mtproto.Phone_PhoneCall, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PhoneConfirmCall - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PhoneConfirmCall(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PhoneConfirmCall - reply: {%s}", r)
	return r, err
}
// PhoneReceivedCall
func (s *Service) PhoneReceivedCall(ctx context.Context, request *mtproto.TLPhoneReceivedCall) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PhoneReceivedCall - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PhoneReceivedCall(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PhoneReceivedCall - reply: {%s}", r)
	return r, err
}
// PhoneDiscardCall
func (s *Service) PhoneDiscardCall(ctx context.Context, request *mtproto.TLPhoneDiscardCall) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PhoneDiscardCall - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PhoneDiscardCall(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PhoneDiscardCall - reply: {%s}", r)
	return r, err
}
// PhoneSetCallRating
func (s *Service) PhoneSetCallRating(ctx context.Context, request *mtproto.TLPhoneSetCallRating) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PhoneSetCallRating - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PhoneSetCallRating(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PhoneSetCallRating - reply: {%s}", r)
	return r, err
}
// PhoneSaveCallDebug
func (s *Service) PhoneSaveCallDebug(ctx context.Context, request *mtproto.TLPhoneSaveCallDebug) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PhoneSaveCallDebug - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PhoneSaveCallDebug(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PhoneSaveCallDebug - reply: {%s}", r)
	return r, err
}
// PhoneSendSignalingData
func (s *Service) PhoneSendSignalingData(ctx context.Context, request *mtproto.TLPhoneSendSignalingData) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PhoneSendSignalingData - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PhoneSendSignalingData(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PhoneSendSignalingData - reply: {%s}", r)
	return r, err
}
// PhoneSaveCallLog
func (s *Service) PhoneSaveCallLog(ctx context.Context, request *mtproto.TLPhoneSaveCallLog) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PhoneSaveCallLog - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PhoneSaveCallLog(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PhoneSaveCallLog - reply: {%s}", r)
	return r, err
}
