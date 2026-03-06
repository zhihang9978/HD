package svc

import (
	"github.com/teamgram/teamgram-server/app/bff/stickers/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
