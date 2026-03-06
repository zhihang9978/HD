package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// FoldersEditPeerFolders
func (c *FoldersCore) FoldersEditPeerFolders(in *mtproto.TLFoldersEditPeerFolders) (*mtproto.Updates, error) {
	c.Logger.Errorf("FoldersEditPeerFolders - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
