package core

import "github.com/teamgram/proto/mtproto"

func (c *DraftsCore) MessagesClearAllDrafts(in *mtproto.TLMessagesClearAllDrafts) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
