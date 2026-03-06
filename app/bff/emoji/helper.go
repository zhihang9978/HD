package emoji_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/emoji/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/emoji/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/emoji/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
