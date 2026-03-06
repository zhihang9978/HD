package core

import (
	"github.com/teamgram/proto/mtproto"
)

// PaymentsGetSavedInfo
func (c *PaymentsCore) PaymentsGetSavedInfo(in *mtproto.TLPaymentsGetSavedInfo) (*mtproto.Payments_SavedInfo, error) {
	c.Logger.Errorf("PaymentsGetSavedInfo - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
