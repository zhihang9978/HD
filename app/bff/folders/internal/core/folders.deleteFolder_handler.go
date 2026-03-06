package core

import (
	"time"
	"github.com/teamgram/proto/mtproto"
)

// FoldersDeleteFolder
func (c *FoldersCore) FoldersDeleteFolder(in *mtproto.TLFoldersDeleteFolder) (*mtproto.Updates, error) {
	c.Logger.Errorf("FoldersDeleteFolder - not impl")

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates(), nil
}
