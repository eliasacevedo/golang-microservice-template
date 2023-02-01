package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/eliasacevedo/golang-microservice-template/config"
	docs "github.com/eliasacevedo/golang-microservice-template/docs"
	"github.com/eliasacevedo/golang-microservice-template/middlewares"
	"github.com/eliasacevedo/golang-microservice-template/server"
	"github.com/eliasacevedo/golang-microservice-template/utilities"
	e "github.com/eliasacevedo/golang-microservice-template/x"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// Base logger
	var l = utilities.NewLogger()

	// Loading environment variables from external file
	config.LoadEnvFromFile(l)

	gin.SetMode(config.GetAppMode())
	modules := e.GetModules()
	handler := server.NewRouter(&l)

	handler.Use(middlewares.DataReturnMiddleware(&l))
	handler.Use(middlewares.EventsMiddleware(&l, config.GetMustLogInfo(), config.GetMustLogValidationError(), config.GetMustLogServerError()))
	for _, module := range modules {
		module.SetRoutes(handler, &l)
	}

	// swagger
	configureSwagger()
	handler.GET(server.SwaggerRoute, ginSwagger.WrapHandler(files.Handler))

	// Running server
	srv := server.NewServer(server.GetServerConfigFromEnvVar(), &l, handler)
	go server.RunServer(srv, l)

	// Detect when server will shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	sig := <-c
	server.OnShutDownServer(srv, l, sig)
}

func configureSwagger() {
	version := config.GetVersion()
	docs.SwaggerInfo.Title = fmt.Sprintf("%s API", config.GetAppName())
	docs.SwaggerInfo.Version = strconv.FormatInt(version, 10)
	docs.SwaggerInfo.Description = config.GetDescription()
	docs.SwaggerInfo.BasePath = fmt.Sprintf("/%s/v%d", config.GetRoutesPrefix(), version)
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
