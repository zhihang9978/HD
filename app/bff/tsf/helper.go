package tsf_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/tsf/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/tsf/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/tsf/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
