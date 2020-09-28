package config

import (
	chi "github.com/common-go/chi-log"
	"github.com/common-go/log"
)

type Root struct {
	Server        ServerConfig `mapstructure:"server"`
	Sql           SqlConfig           `mapstructure:"sql"`
	Log           log.Config          `mapstructure:"log"`
	MiddleWareLog chi.ChiLogConfig    `mapstructure:"middleware_log"`
}

type ServerConfig struct {
	Name       string `mapstructure:"name"`
	Version    string `mapstructure:"version"`
	Port       int    `mapstructure:"port"`
}

type SqlConfig struct {
	Uri string `mapstructure:"uri"`
}
