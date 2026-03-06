package core

import (
	"github.com/teamgram/proto/mtproto"
)

// MessagesGetPollVotes
func (c *PollsCore) MessagesGetPollVotes(in *mtproto.TLMessagesGetPollVotes) (*mtproto.Messages_VotesList, error) {
	c.Logger.Errorf("MessagesGetPollVotes - not impl")

	return nil, mtproto.ErrMethodNotImpl
}
