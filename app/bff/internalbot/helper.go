package internalbot_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/internalbot/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/internalbot/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/internalbot/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
