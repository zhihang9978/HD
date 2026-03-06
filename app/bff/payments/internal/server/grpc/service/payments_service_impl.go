package service

import (
	"context"

	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/bff/payments/internal/core"
)

// AccountGetTmpPassword
func (s *Service) AccountGetTmpPassword(ctx context.Context, request *mtproto.TLAccountGetTmpPassword) (*mtproto.Account_TmpPassword, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("AccountGetTmpPassword - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.AccountGetTmpPassword(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("AccountGetTmpPassword - reply: {%s}", r)
	return r, err
}
// MessagesSetBotShippingResults
func (s *Service) MessagesSetBotShippingResults(ctx context.Context, request *mtproto.TLMessagesSetBotShippingResults) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSetBotShippingResults - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSetBotShippingResults(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSetBotShippingResults - reply: {%s}", r)
	return r, err
}
// MessagesSetBotPrecheckoutResults
func (s *Service) MessagesSetBotPrecheckoutResults(ctx context.Context, request *mtproto.TLMessagesSetBotPrecheckoutResults) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("MessagesSetBotPrecheckoutResults - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.MessagesSetBotPrecheckoutResults(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("MessagesSetBotPrecheckoutResults - reply: {%s}", r)
	return r, err
}
// PaymentsGetPaymentForm
func (s *Service) PaymentsGetPaymentForm(ctx context.Context, request *mtproto.TLPaymentsGetPaymentForm) (*mtproto.Payments_PaymentForm, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PaymentsGetPaymentForm - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PaymentsGetPaymentForm(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PaymentsGetPaymentForm - reply: {%s}", r)
	return r, err
}
// PaymentsGetPaymentReceipt
func (s *Service) PaymentsGetPaymentReceipt(ctx context.Context, request *mtproto.TLPaymentsGetPaymentReceipt) (*mtproto.Payments_PaymentReceipt, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PaymentsGetPaymentReceipt - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PaymentsGetPaymentReceipt(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PaymentsGetPaymentReceipt - reply: {%s}", r)
	return r, err
}
// PaymentsValidateRequestedInfo
func (s *Service) PaymentsValidateRequestedInfo(ctx context.Context, request *mtproto.TLPaymentsValidateRequestedInfo) (*mtproto.Payments_ValidatedRequestedInfo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PaymentsValidateRequestedInfo - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PaymentsValidateRequestedInfo(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PaymentsValidateRequestedInfo - reply: {%s}", r)
	return r, err
}
// PaymentsSendPaymentForm
func (s *Service) PaymentsSendPaymentForm(ctx context.Context, request *mtproto.TLPaymentsSendPaymentForm) (*mtproto.Payments_PaymentResult, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PaymentsSendPaymentForm - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PaymentsSendPaymentForm(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PaymentsSendPaymentForm - reply: {%s}", r)
	return r, err
}
// PaymentsGetSavedInfo
func (s *Service) PaymentsGetSavedInfo(ctx context.Context, request *mtproto.TLPaymentsGetSavedInfo) (*mtproto.Payments_SavedInfo, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PaymentsGetSavedInfo - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PaymentsGetSavedInfo(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PaymentsGetSavedInfo - reply: {%s}", r)
	return r, err
}
// PaymentsClearSavedInfo
func (s *Service) PaymentsClearSavedInfo(ctx context.Context, request *mtproto.TLPaymentsClearSavedInfo) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PaymentsClearSavedInfo - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PaymentsClearSavedInfo(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PaymentsClearSavedInfo - reply: {%s}", r)
	return r, err
}
// PaymentsGetBankCardData
func (s *Service) PaymentsGetBankCardData(ctx context.Context, request *mtproto.TLPaymentsGetBankCardData) (*mtproto.Payments_BankCardData, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PaymentsGetBankCardData - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PaymentsGetBankCardData(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PaymentsGetBankCardData - reply: {%s}", r)
	return r, err
}
// PaymentsExportInvoice
func (s *Service) PaymentsExportInvoice(ctx context.Context, request *mtproto.TLPaymentsExportInvoice) (*mtproto.Payments_ExportedInvoice, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PaymentsExportInvoice - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PaymentsExportInvoice(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PaymentsExportInvoice - reply: {%s}", r)
	return r, err
}
// PaymentsRequestRecurringPayment
func (s *Service) PaymentsRequestRecurringPayment(ctx context.Context, request *mtproto.TLPaymentsRequestRecurringPayment) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PaymentsRequestRecurringPayment - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PaymentsRequestRecurringPayment(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PaymentsRequestRecurringPayment - reply: {%s}", r)
	return r, err
}
// PaymentsRestorePlayMarketReceipt
func (s *Service) PaymentsRestorePlayMarketReceipt(ctx context.Context, request *mtproto.TLPaymentsRestorePlayMarketReceipt) (*mtproto.Updates, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("PaymentsRestorePlayMarketReceipt - metadata: {%s}, request: {%s}", c.MD, request)

	r, err := c.PaymentsRestorePlayMarketReceipt(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("PaymentsRestorePlayMarketReceipt - reply: {%s}", r)
	return r, err
}
