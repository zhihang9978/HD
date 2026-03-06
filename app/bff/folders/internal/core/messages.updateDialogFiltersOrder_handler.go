package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesUpdateDialogFiltersOrder
func (c *FoldersCore) MessagesUpdateDialogFiltersOrder(in *mtproto.TLMessagesUpdateDialogFiltersOrder) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesUpdateDialogFiltersOrder - not impl")

	return mtproto.BoolTrue, nil
}
