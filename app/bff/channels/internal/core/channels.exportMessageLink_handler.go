package core

import (
	"github.com/teamgram/proto/mtproto"
)

// ChannelsExportMessageLink
func (c *ChannelsCore) ChannelsExportMessageLink(in *mtproto.TLChannelsExportMessageLink) (*mtproto.ExportedMessageLink, error) {
	c.Logger.Errorf("ChannelsExportMessageLink - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
