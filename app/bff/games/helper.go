package games_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/games/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/games/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/games/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
