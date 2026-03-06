package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesInstallStickerSet
func (c *StickersCore) MessagesInstallStickerSet(in *mtproto.TLMessagesInstallStickerSet) (*mtproto.Messages_StickerSetInstallResult, error) {
	c.Logger.Errorf("MessagesInstallStickerSet - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
