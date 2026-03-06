package core

import "github.com/teamgram/proto/mtproto"

func (c *MiscellaneousCore) HelpTest(in *mtproto.TLHelpTest) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
