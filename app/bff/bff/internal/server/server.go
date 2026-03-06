// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package server

import (
	"flag"

	"github.com/teamgram/proto/mtproto"
	account_helper "github.com/teamgram/teamgram-server/app/bff/account"
	authorization_helper "github.com/teamgram/teamgram-server/app/bff/authorization"
	autodownload_helper "github.com/teamgram/teamgram-server/app/bff/autodownload"
	"github.com/teamgram/teamgram-server/app/bff/bff/internal/config"
	bots_helper "github.com/teamgram/teamgram-server/app/bff/bots"
	channels_helper "github.com/teamgram/teamgram-server/app/bff/channels"
	chatinvites_helper "github.com/teamgram/teamgram-server/app/bff/chatinvites"
	chats_helper "github.com/teamgram/teamgram-server/app/bff/chats"
	configuration_helper "github.com/teamgram/teamgram-server/app/bff/configuration"
	contacts_helper "github.com/teamgram/teamgram-server/app/bff/contacts"
	deeplinks_helper "github.com/teamgram/teamgram-server/app/bff/deeplinks"
	dialogs_helper "github.com/teamgram/teamgram-server/app/bff/dialogs"
	drafts_helper "github.com/teamgram/teamgram-server/app/bff/drafts"
	emoji_helper "github.com/teamgram/teamgram-server/app/bff/emoji"
	files_helper "github.com/teamgram/teamgram-server/app/bff/files"
	folders_helper "github.com/teamgram/teamgram-server/app/bff/folders"
	games_helper "github.com/teamgram/teamgram-server/app/bff/games"
	gifs_helper "github.com/teamgram/teamgram-server/app/bff/gifs"
	groupcalls_helper "github.com/teamgram/teamgram-server/app/bff/groupcalls"
	importedchats_helper "github.com/teamgram/teamgram-server/app/bff/importedchats"
	inlinebot_helper "github.com/teamgram/teamgram-server/app/bff/inlinebot"
	internalbot_helper "github.com/teamgram/teamgram-server/app/bff/internalbot"
	langpack_helper "github.com/teamgram/teamgram-server/app/bff/langpack"
	messages_helper "github.com/teamgram/teamgram-server/app/bff/messages"
	messagethreads_helper "github.com/teamgram/teamgram-server/app/bff/messagethreads"
	miscellaneous_helper "github.com/teamgram/teamgram-server/app/bff/miscellaneous"
	notification_helper "github.com/teamgram/teamgram-server/app/bff/notification"
	nsfw_helper "github.com/teamgram/teamgram-server/app/bff/nsfw"
	passkeyhelper "github.com/teamgram/teamgram-server/app/bff/passkey"
	passport_helper "github.com/teamgram/teamgram-server/app/bff/passport"
	payments_helper "github.com/teamgram/teamgram-server/app/bff/payments"
	polls_helper "github.com/teamgram/teamgram-server/app/bff/polls"
	premium_helper "github.com/teamgram/teamgram-server/app/bff/premium"
	privacysettingshelper "github.com/teamgram/teamgram-server/app/bff/privacysettings"
	promodata_helper "github.com/teamgram/teamgram-server/app/bff/promodata"
	qrcode_helper "github.com/teamgram/teamgram-server/app/bff/qrcode"
	reactions_helper "github.com/teamgram/teamgram-server/app/bff/reactions"
	reports_helper "github.com/teamgram/teamgram-server/app/bff/reports"
	savedmessagedialogshelper "github.com/teamgram/teamgram-server/app/bff/savedmessagedialogs"
	scheduledmessages_helper "github.com/teamgram/teamgram-server/app/bff/scheduledmessages"
	seamless_helper "github.com/teamgram/teamgram-server/app/bff/seamless"
	secretchats_helper "github.com/teamgram/teamgram-server/app/bff/secretchats"
	sponsoredmessages_helper "github.com/teamgram/teamgram-server/app/bff/sponsoredmessages"
	statistics_helper "github.com/teamgram/teamgram-server/app/bff/statistics"
	stickers_helper "github.com/teamgram/teamgram-server/app/bff/stickers"
	themes_helper "github.com/teamgram/teamgram-server/app/bff/themes"
	tos_helper "github.com/teamgram/teamgram-server/app/bff/tos"
	tsf_helper "github.com/teamgram/teamgram-server/app/bff/tsf"
	twofa_helper "github.com/teamgram/teamgram-server/app/bff/twofa"
	updates_helper "github.com/teamgram/teamgram-server/app/bff/updates"
	userchannelprofileshelper "github.com/teamgram/teamgram-server/app/bff/userchannelprofiles"
	usernames_helper "github.com/teamgram/teamgram-server/app/bff/usernames"
	users_helper "github.com/teamgram/teamgram-server/app/bff/users"
	voipcalls_helper "github.com/teamgram/teamgram-server/app/bff/voipcalls"
	wallpapers_helper "github.com/teamgram/teamgram-server/app/bff/wallpapers"
	webpage_helper "github.com/teamgram/teamgram-server/app/bff/webpage"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/bff.yaml", "the config file")

type Server struct {
	grpcSrv *zrpc.RpcServer
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() error {
	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.Infov(c)
	// ctx := svc.NewServiceContext(c)
	// s.grpcSrv = grpc.New(ctx, c.RpcServerConf)

	s.grpcSrv = zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		// tos_helper
		mtproto.RegisterRPCTosServer(
			grpcServer,
			tos_helper.New(tos_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// configuration_helper
		mtproto.RegisterRPCConfigurationServer(
			grpcServer,
			configuration_helper.New(configuration_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// qrcode_helper
		mtproto.RegisterRPCQrCodeServer(
			grpcServer,
			qrcode_helper.New(
				qrcode_helper.Config{
					RpcServerConf:     c.RpcServerConf,
					KV:                c.KV,
					UserClient:        c.BizServiceClient,
					AuthSessionClient: c.AuthSessionClient,
					SyncClient:        c.SyncClient,
				},
				nil))

		// miscellaneous_helper
		mtproto.RegisterRPCMiscellaneousServer(
			grpcServer,
			miscellaneous_helper.New(miscellaneous_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// authorization_helper
		mtproto.RegisterRPCAuthorizationServer(
			grpcServer,
			authorization_helper.New(
				authorization_helper.Config{
					RpcServerConf:             c.RpcServerConf,
					KV:                        c.KV,
					Code:                      c.Code,
					UserClient:                c.BizServiceClient,
					AuthsessionClient:         c.AuthSessionClient,
					ChatClient:                c.BizServiceClient,
					StatusClient:              c.StatusClient,
					SyncClient:                c.SyncClient,
					MsgClient:                 c.MsgClient,
					SignInMessage:             c.SignInMessage,
					SignInServiceNotification: c.SignInServiceNotification,
				},
				nil,
				nil))

		// premium_helper
		mtproto.RegisterRPCPremiumServer(
			grpcServer,
			premium_helper.New(premium_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// chatinvites_helper
		mtproto.RegisterRPCChatInvitesServer(
			grpcServer,
			chatinvites_helper.New(chatinvites_helper.Config{
				RpcServerConf: c.RpcServerConf,
				UserClient:    c.BizServiceClient,
				ChatClient:    c.BizServiceClient,
				MsgClient:     c.MsgClient,
				SyncClient:    c.SyncClient,
			}))

		// chats_helper
		mtproto.RegisterRPCChatsServer(
			grpcServer,
			chats_helper.New(chats_helper.Config{
				RpcServerConf:     c.RpcServerConf,
				UserClient:        c.BizServiceClient,
				ChatClient:        c.BizServiceClient,
				MsgClient:         c.MsgClient,
				DialogClient:      c.BizServiceClient,
				SyncClient:        c.SyncClient,
				MediaClient:       c.MediaClient,
				AuthsessionClient: c.AuthSessionClient,
				IdgenClient:       c.IdgenClient,
				MessageClient:     c.BizServiceClient,
			}))

		// files_helper
		mtproto.RegisterRPCFilesServer(
			grpcServer,
			files_helper.New(files_helper.Config{
				RpcServerConf: c.RpcServerConf,
				DfsClient:     c.DfsClient,
				UserClient:    c.BizServiceClient,
				MediaClient:   c.MediaClient,
			}, nil, nil))

		// passport_helper
		mtproto.RegisterRPCPassportServer(
			grpcServer,
			passport_helper.New(passport_helper.Config{
				RpcServerConf:     c.RpcServerConf,
				AuthsessionClient: c.AuthSessionClient,
				UserClient:        c.BizServiceClient,
			}))

		// updates_helper
		mtproto.RegisterRPCUpdatesServer(
			grpcServer,
			updates_helper.New(updates_helper.Config{
				RpcServerConf:     c.RpcServerConf,
				UpdatesClient:     c.BizServiceClient,
				UserClient:        c.BizServiceClient,
				ChatClient:        c.BizServiceClient,
				AuthsessionClient: c.AuthSessionClient,
			}))

		// contacts_helper
		mtproto.RegisterRPCContactsServer(
			grpcServer,
			contacts_helper.New(
				contacts_helper.Config{
					RpcServerConf: c.RpcServerConf,
					UserClient:    c.BizServiceClient,
					ChatClient:    c.BizServiceClient,
					SyncClient:    c.SyncClient,
				},
				nil))

		// dialogs_helper
		mtproto.RegisterRPCDialogsServer(
			grpcServer,
			dialogs_helper.New(dialogs_helper.Config{
				RpcServerConf: c.RpcServerConf,
				UpdatesClient: c.BizServiceClient,
				UserClient:    c.BizServiceClient,
				ChatClient:    c.BizServiceClient,
				DialogClient:  c.BizServiceClient,
				SyncClient:    c.SyncClient,
				MessageClient: c.BizServiceClient,
			}, nil))

		// drafts_helper
		mtproto.RegisterRPCDraftsServer(
			grpcServer,
			drafts_helper.New(drafts_helper.Config{
				RpcServerConf: c.RpcServerConf,
				DialogClient:  c.BizServiceClient,
				UserClient:    c.BizServiceClient,
				SyncClient:    c.SyncClient,
				ChatClient:    c.BizServiceClient,
			}, nil))

		// autodownload_helper
		mtproto.RegisterRPCAutoDownloadServer(
			grpcServer,
			autodownload_helper.New(autodownload_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// messages_helper
		mtproto.RegisterRPCMessagesServer(
			grpcServer,
			messages_helper.New(messages_helper.Config{
				RpcServerConf: c.RpcServerConf,
				UserClient:    c.BizServiceClient,
				ChatClient:    c.BizServiceClient,
				MsgClient:     c.MsgClient,
				DialogClient:  c.BizServiceClient,
				IdgenClient:   c.IdgenClient,
				MessageClient: c.BizServiceClient,
				MediaClient:   c.MediaClient,
				SyncClient:    c.SyncClient,
			}, nil))

		// notification_helper
		mtproto.RegisterRPCNotificationServer(
			grpcServer,
			notification_helper.New(notification_helper.Config{
				RpcServerConf: c.RpcServerConf,
				UserClient:    c.BizServiceClient,
				ChatClient:    c.BizServiceClient,
				SyncClient:    c.SyncClient,
			}, nil))

		// users_helper
		mtproto.RegisterRPCUsersServer(
			grpcServer,
			users_helper.New(
				users_helper.Config{
					RpcServerConf: c.RpcServerConf,
					UserClient:    c.BizServiceClient,
					ChatClient:    c.BizServiceClient,
					DialogClient:  c.BizServiceClient,
				},
				nil,
				nil,
				nil))

		// nsfw_helper
		mtproto.RegisterRPCNsfwServer(
			grpcServer,
			nsfw_helper.New(nsfw_helper.Config{
				RpcServerConf: c.RpcServerConf,
				UserClient:    c.BizServiceClient,
			}))

		// sponsoredmessages_helper
		mtproto.RegisterRPCSponsoredMessagesServer(
			grpcServer,
			sponsoredmessages_helper.New(sponsoredmessages_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// account_helper
		mtproto.RegisterRPCAccountServer(
			grpcServer,
			account_helper.New(
				account_helper.Config{
					RpcServerConf:     c.RpcServerConf,
					KV:                c.KV,
					UserClient:        c.BizServiceClient,
					AuthsessionClient: c.AuthSessionClient,
					ChatClient:        c.BizServiceClient,
					SyncClient:        c.SyncClient,
				},
				nil,
				nil))
		// usernames_helper
		mtproto.RegisterRPCUsernamesServer(
			grpcServer,
			usernames_helper.New(usernames_helper.Config{
				RpcServerConf: c.RpcServerConf,
				UserClient:    c.BizServiceClient,
				ChatClient:    c.BizServiceClient,
				SyncClient:    c.SyncClient,
			}, nil))

		// privacysettingshelper
		mtproto.RegisterRPCPrivacySettingsServer(
			grpcServer,
			privacysettingshelper.New(privacysettingshelper.Config{
				RpcServerConf:     c.RpcServerConf,
				UserClient:        c.BizServiceClient,
				AuthsessionClient: c.AuthSessionClient,
				ChatClient:        c.BizServiceClient,
				SyncClient:        c.SyncClient,
			}))

		// savedmessagedialogshelper
		mtproto.RegisterRPCSavedMessageDialogsServer(
			grpcServer,
			savedmessagedialogshelper.New(savedmessagedialogshelper.Config{
				RpcServerConf: c.RpcServerConf,
				UpdatesClient: c.BizServiceClient,
				UserClient:    c.BizServiceClient,
				ChatClient:    c.BizServiceClient,
				DialogClient:  c.BizServiceClient,
				SyncClient:    c.SyncClient,
				MessageClient: c.BizServiceClient,
			}))

		// userchannelprofileshelper
		mtproto.RegisterRPCUserChannelProfilesServer(
			grpcServer,
			userchannelprofileshelper.New(userchannelprofileshelper.Config{
				RpcServerConf: c.RpcServerConf,
				MediaClient:   c.MediaClient,
				UserClient:    c.BizServiceClient,
				SyncClient:    c.SyncClient,
			}))

		mtproto.RegisterRPCPasskeyServer(
			grpcServer,
			passkeyhelper.New(passkeyhelper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// reports_helper
		mtproto.RegisterRPCReportsServer(
			grpcServer,
			reports_helper.New(reports_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// gifs_helper
		mtproto.RegisterRPCGifsServer(
			grpcServer,
			gifs_helper.New(gifs_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// promodata_helper
		mtproto.RegisterRPCPromoDataServer(
			grpcServer,
			promodata_helper.New(promodata_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// tsf_helper
		mtproto.RegisterRPCTsfServer(
			grpcServer,
			tsf_helper.New(tsf_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// twofa_helper
		mtproto.RegisterRPCTwoFaServer(
			grpcServer,
			twofa_helper.New(twofa_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// seamless_helper
		mtproto.RegisterRPCSeamlessServer(
			grpcServer,
			seamless_helper.New(seamless_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// voipcalls_helper
		mtproto.RegisterRPCVoipCallsServer(
			grpcServer,
			voipcalls_helper.New(voipcalls_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// channels_helper
		mtproto.RegisterRPCChannelsServer(
			grpcServer,
			channels_helper.New(channels_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// deeplinks_helper
		mtproto.RegisterRPCDeepLinksServer(
			grpcServer,
			deeplinks_helper.New(deeplinks_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// webpage_helper
		mtproto.RegisterRPCWebPageServer(
			grpcServer,
			webpage_helper.New(webpage_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// secretchats_helper
		mtproto.RegisterRPCSecretChatsServer(
			grpcServer,
			secretchats_helper.New(secretchats_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// inlinebot_helper
		mtproto.RegisterRPCInlineBotServer(
			grpcServer,
			inlinebot_helper.New(inlinebot_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// bots_helper
		mtproto.RegisterRPCBotsServer(
			grpcServer,
			bots_helper.New(bots_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// internalbot_helper
		mtproto.RegisterRPCInternalBotServer(
			grpcServer,
			internalbot_helper.New(internalbot_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// themes_helper
		mtproto.RegisterRPCThemesServer(
			grpcServer,
			themes_helper.New(themes_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// emoji_helper
		mtproto.RegisterRPCEmojiServer(
			grpcServer,
			emoji_helper.New(emoji_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// folders_helper
		mtproto.RegisterRPCFoldersServer(
			grpcServer,
			folders_helper.New(folders_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// games_helper
		mtproto.RegisterRPCGamesServer(
			grpcServer,
			games_helper.New(games_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// groupcalls_helper
		mtproto.RegisterRPCGroupCallsServer(
			grpcServer,
			groupcalls_helper.New(groupcalls_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// importedchats_helper
		mtproto.RegisterRPCImportedChatsServer(
			grpcServer,
			importedchats_helper.New(importedchats_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// langpack_helper
		mtproto.RegisterRPCLangpackServer(
			grpcServer,
			langpack_helper.New(langpack_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// messagethreads_helper
		mtproto.RegisterRPCMessageThreadsServer(
			grpcServer,
			messagethreads_helper.New(messagethreads_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// reactions_helper
		mtproto.RegisterRPCReactionsServer(
			grpcServer,
			reactions_helper.New(reactions_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// payments_helper
		mtproto.RegisterRPCPaymentsServer(
			grpcServer,
			payments_helper.New(payments_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// polls_helper
		mtproto.RegisterRPCPollsServer(
			grpcServer,
			polls_helper.New(polls_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// scheduledmessages_helper
		mtproto.RegisterRPCScheduledMessagesServer(
			grpcServer,
			scheduledmessages_helper.New(scheduledmessages_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// statistics_helper
		mtproto.RegisterRPCStatisticsServer(
			grpcServer,
			statistics_helper.New(statistics_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// stickers_helper
		mtproto.RegisterRPCStickersServer(
			grpcServer,
			stickers_helper.New(stickers_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))

		// wallpapers_helper
		mtproto.RegisterRPCWallpapersServer(
			grpcServer,
			wallpapers_helper.New(wallpapers_helper.Config{
				RpcServerConf: c.RpcServerConf,
			}))
	})

	// logx.Must(err)

	go func() {
		s.grpcSrv.Start()
	}()
	return nil
}

func (s *Server) RunLoop() {
}

func (s *Server) Destroy() {
	s.grpcSrv.Stop()
}
