package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PaymentsExportInvoice
func (c *PaymentsCore) PaymentsExportInvoice(in *mtproto.TLPaymentsExportInvoice) (*mtproto.Payments_ExportedInvoice, error) {
	c.Logger.Errorf("PaymentsExportInvoice - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
