package seamless_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/seamless/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/seamless/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/seamless/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
