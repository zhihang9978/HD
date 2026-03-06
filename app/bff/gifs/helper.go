package gifs_helper

import (
	"github.com/teamgram/teamgram-server/app/bff/gifs/internal/config"
	"github.com/teamgram/teamgram-server/app/bff/gifs/internal/server/grpc/service"
	"github.com/teamgram/teamgram-server/app/bff/gifs/internal/svc"
)

type (
	Config = config.Config
)

func New(c Config) *service.Service {
	return service.New(svc.NewServiceContext(c))
}
