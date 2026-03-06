package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetDialogFiltersF19ED96D
func (c *FoldersCore) MessagesGetDialogFiltersF19ED96D(in *mtproto.TLMessagesGetDialogFiltersF19ED96D) (*mtproto.Vector_DialogFilter, error) {
	c.Logger.Errorf("MessagesGetDialogFiltersF19ED96D - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
