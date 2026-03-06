package core

import (
	"github.com/teamgram/proto/mtproto"
)

// AccountReportProfilePhoto
func (c *ReportsCore) AccountReportProfilePhoto(in *mtproto.TLAccountReportProfilePhoto) (*mtproto.Bool, error) {
	c.Logger.Errorf("AccountReportProfilePhoto - not impl")

	return mtproto.BoolTrue, nil
}
