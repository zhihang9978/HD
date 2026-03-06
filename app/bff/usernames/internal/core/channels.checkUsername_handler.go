package core

import "github.com/teamgram/proto/mtproto"

func (c *UsernamesCore) ChannelsCheckUsername(in *mtproto.TLChannelsCheckUsername) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
