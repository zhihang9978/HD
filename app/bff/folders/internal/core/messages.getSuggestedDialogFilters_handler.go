package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetSuggestedDialogFilters
func (c *FoldersCore) MessagesGetSuggestedDialogFilters(in *mtproto.TLMessagesGetSuggestedDialogFilters) (*mtproto.Vector_DialogFilterSuggested, error) {
	c.Logger.Errorf("MessagesGetSuggestedDialogFilters - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
