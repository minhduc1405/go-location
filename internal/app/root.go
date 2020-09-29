package app

import (
	chi "github.com/common-go/chi-log"
	"github.com/common-go/log"
	"github.com/common-go/mongo"
	"github.com/common-go/web"
)

type Root struct {
	Server        server.ServerConfig `mapstructure:"server"`
	Mongo         mongo.MongoConfig   `mapstructure:"mongo"`
	Sql           SqlConfig           `mapstructure:"sql"`
	Log           log.Config          `mapstructure:"log"`
	MiddleWareLog chi.ChiLogConfig    `mapstructure:"middleware_log"`
}

type SqlConfig struct {
	Uri string `mapstructure:"uri"`
}
