package core

import "github.com/teamgram/proto/mtproto"

func (c *ConfigurationCore) HelpDismissSuggestion(in *mtproto.TLHelpDismissSuggestion) (*mtproto.Bool, error) {
	return mtproto.BoolTrue, nil
}
