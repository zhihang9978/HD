package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesUpdateDialogFilter
func (c *FoldersCore) MessagesUpdateDialogFilter(in *mtproto.TLMessagesUpdateDialogFilter) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesUpdateDialogFilter - not impl")

	return mtproto.BoolTrue, nil
}
