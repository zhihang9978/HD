package secretchats_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/secretchats/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/secretchats/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/secretchats/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
