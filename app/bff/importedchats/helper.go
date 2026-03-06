package importedchats_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/importedchats/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/importedchats/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/importedchats/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
