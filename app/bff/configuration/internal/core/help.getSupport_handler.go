package core

import (
	"github.com/teamgram/proto/mtproto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (c *ConfigurationCore) HelpGetSupport(in *mtproto.TLHelpGetSupport) (*mtproto.Help_Support, error) {
	return mtproto.MakeTLHelpSupport(&mtproto.Help_Support{
		PhoneNumber: "+86000000000",
		User: mtproto.MakeTLUser(&mtproto.User{
			Id: 777000, Self: false, Contact: false,
			FirstName: &wrapperspb.StringValue{Value: "Support"},
			Username:  &wrapperspb.StringValue{Value: "support"},
			Phone:     &wrapperspb.StringValue{Value: "+86000000000"},
			Status:    mtproto.MakeTLUserStatusRecently(&mtproto.UserStatus{}).To_UserStatus(),
		}).To_User(),
	}).To_Help_Support(), nil
}
