package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetDialogFiltersEFD48C89
func (c *FoldersCore) MessagesGetDialogFiltersEFD48C89(in *mtproto.TLMessagesGetDialogFiltersEFD48C89) (*mtproto.Messages_DialogFilters, error) {
	c.Logger.Errorf("MessagesGetDialogFiltersEFD48C89 - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
