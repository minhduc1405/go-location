package main

import (
	"context"
	"fmt"
	chilog "github.com/common-go/chi-log"
	"github.com/common-go/config"
	"github.com/common-go/log"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"os"
	"strconv"

	c "go-location/internal/app"
)

func main() {
	parentPath := "go-location"
	resource := "configs"
	env := os.Getenv("ENV")

	conf := c.Root{}
	config.LoadConfig(parentPath, resource, env, &conf, "application")

	r := chi.NewRouter()

	logger := log.Initialize(conf.Log)
	logger.SetOutput(os.Stdout)

	r.Use(middleware.RequestID)
	l := chilog.NewStructuredLogger(logger)
	r.Use(chilog.Logger(chilog.Standardize(conf.MiddleWareLog), logger, l))

	er2 := c.Route(r, context.Background(), conf)
	if er2 != nil {
		panic(er2)
	}
	fmt.Println("Start server")
	server := ""
	if conf.Server.Port > 0 {
		server = ":" + strconv.Itoa(conf.Server.Port)
	}
	er3 := http.ListenAndServe(server, r)
	if er3 != nil {
		fmt.Println(er3.Error())
	}
}
