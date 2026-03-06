package core

import "github.com/teamgram/proto/mtproto"

func (c *PrivacySettingsCore) UsersGetIsPremiumRequiredToContact(in *mtproto.TLUsersGetIsPremiumRequiredToContact) (*mtproto.Vector_Bool, error) {
	bools := make([]*mtproto.Bool, len(in.GetId()))
	for i := range in.GetId() {
		bools[i] = mtproto.BoolFalse
	}
	return &mtproto.Vector_Bool{Datas: bools}, nil
}
