package core

import "github.com/teamgram/proto/mtproto"

func (c *MessagesCore) ChannelsGetSendAs(in *mtproto.TLChannelsGetSendAs) (*mtproto.Channels_SendAsPeers, error) {
	return mtproto.MakeTLChannelsSendAsPeers(&mtproto.Channels_SendAsPeers{
		Peers_VECTORSENDASPEER: []*mtproto.SendAsPeer{},
		Chats: []*mtproto.Chat{},
		Users: []*mtproto.User{},
	}).To_Channels_SendAsPeers(), nil
}
