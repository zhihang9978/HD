package core

import "github.com/teamgram/proto/mtproto"

func (c *DraftsCore) MessagesSaveDraft(in *mtproto.TLMessagesSaveDraft) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
