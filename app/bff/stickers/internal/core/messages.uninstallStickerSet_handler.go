package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesUninstallStickerSet
func (c *StickersCore) MessagesUninstallStickerSet(in *mtproto.TLMessagesUninstallStickerSet) (*mtproto.Bool, error) {
	c.Logger.Errorf("MessagesUninstallStickerSet - not impl")

	return mtproto.BoolTrue, nil
}
