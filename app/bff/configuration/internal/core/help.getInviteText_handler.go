package core

import "github.com/teamgram/proto/mtproto"

func (c *ConfigurationCore) HelpGetInviteText(in *mtproto.TLHelpGetInviteText) (*mtproto.Help_InviteText, error) {
	return mtproto.MakeTLHelpInviteText(&mtproto.Help_InviteText{
		Message: "Hey, I am using PulseChat. Join me! Download it here: https://pulsechat.app",
	}).To_Help_InviteText(), nil
}
