package core

import "github.com/teamgram/proto/mtproto"

func (c *UsernamesCore) ChannelsUpdateUsername(in *mtproto.TLChannelsUpdateUsername) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
