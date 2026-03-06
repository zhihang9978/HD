package core

import "github.com/teamgram/proto/mtproto"

func (c *AutoDownloadCore) AccountSaveAutoDownloadSettings(in *mtproto.TLAccountSaveAutoDownloadSettings) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
