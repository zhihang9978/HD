package core

import "github.com/teamgram/proto/mtproto"

func (c *FilesCore) MessagesGetDocumentByHash(in *mtproto.TLMessagesGetDocumentByHash) (*mtproto.Document, error) {
	return mtproto.MakeTLDocumentEmpty(&mtproto.Document{Id: 0}).To_Document(), nil
}
