package scheduledmessages_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/scheduledmessages/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/scheduledmessages/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/scheduledmessages/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
