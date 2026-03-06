package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/secretchats/internal/core"
)

// MessagesGetDhConfig
func (s *Service) MessagesGetDhConfig(ctx context.Context, request *mtproto.TLMessagesGetDhConfig) (*mtproto.Messages_DhConfig, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesGetDhConfig - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesGetDhConfig(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesGetDhConfig - reply: {%s}", r)
	return r, err
}
// MessagesRequestEncryption
func (s *Service) MessagesRequestEncryption(ctx context.Context, request *mtproto.TLMessagesRequestEncryption) (*mtproto.EncryptedChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesRequestEncryption - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesRequestEncryption(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesRequestEncryption - reply: {%s}", r)
	return r, err
}
// MessagesAcceptEncryption
func (s *Service) MessagesAcceptEncryption(ctx context.Context, request *mtproto.TLMessagesAcceptEncryption) (*mtproto.EncryptedChat, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesAcceptEncryption - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesAcceptEncryption(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesAcceptEncryption - reply: {%s}", r)
	return r, err
}
// MessagesDiscardEncryption
func (s *Service) MessagesDiscardEncryption(ctx context.Context, request *mtproto.TLMessagesDiscardEncryption) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesDiscardEncryption - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesDiscardEncryption(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesDiscardEncryption - reply: {%s}", r)
	return r, err
}
// MessagesSetEncryptedTyping
func (s *Service) MessagesSetEncryptedTyping(ctx context.Context, request *mtproto.TLMessagesSetEncryptedTyping) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSetEncryptedTyping - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSetEncryptedTyping(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSetEncryptedTyping - reply: {%s}", r)
	return r, err
}
// MessagesReadEncryptedHistory
func (s *Service) MessagesReadEncryptedHistory(ctx context.Context, request *mtproto.TLMessagesReadEncryptedHistory) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesReadEncryptedHistory - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesReadEncryptedHistory(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesReadEncryptedHistory - reply: {%s}", r)
	return r, err
}
// MessagesSendEncrypted
func (s *Service) MessagesSendEncrypted(ctx context.Context, request *mtproto.TLMessagesSendEncrypted) (*mtproto.Messages_SentEncryptedMessage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSendEncrypted - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSendEncrypted(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSendEncrypted - reply: {%s}", r)
	return r, err
}
// MessagesSendEncryptedFile
func (s *Service) MessagesSendEncryptedFile(ctx context.Context, request *mtproto.TLMessagesSendEncryptedFile) (*mtproto.Messages_SentEncryptedMessage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSendEncryptedFile - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSendEncryptedFile(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSendEncryptedFile - reply: {%s}", r)
	return r, err
}
// MessagesSendEncryptedService
func (s *Service) MessagesSendEncryptedService(ctx context.Context, request *mtproto.TLMessagesSendEncryptedService) (*mtproto.Messages_SentEncryptedMessage, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSendEncryptedService - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSendEncryptedService(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSendEncryptedService - reply: {%s}", r)
	return r, err
}
// MessagesReceivedQueue
func (s *Service) MessagesReceivedQueue(ctx context.Context, request *mtproto.TLMessagesReceivedQueue) (*mtproto.Vector_Long, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesReceivedQueue - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesReceivedQueue(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesReceivedQueue - reply: {%s}", r)
	return r, err
}
