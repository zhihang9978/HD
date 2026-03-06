package core

import "github.com/teamgram/proto/mtproto"

func (c *MiscellaneousCore) HelpSaveAppLog(in *mtproto.TLHelpSaveAppLog) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
