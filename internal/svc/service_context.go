package svc

import (
	"github.com/xh-polaris/meowchat-moment-rpc/internal/config"
	"github.com/xh-polaris/meowchat-moment-rpc/internal/model"
)

type ServiceContext struct {
	Config      config.Config
	MomentModel model.MomentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		MomentModel: model.NewMomentModel(c.Mongo.URL, c.Mongo.DB, c.Cache),
	}
}
