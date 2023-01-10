package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/eliasacevedo/golang-microservice-template/config"
	"github.com/eliasacevedo/golang-microservice-template/middlewares"
	"github.com/eliasacevedo/golang-microservice-template/server"
	"github.com/eliasacevedo/golang-microservice-template/utilities"
	e "github.com/eliasacevedo/golang-microservice-template/x"
)

func main() {
	// Base logger
	var l = utilities.NewLogger()

	// Loading environment variables from external file
	config.LoadEnvFromFile(l)

	modules := e.GetModules()
	handler := server.NewRouter(&l)
	handler.Use(middlewares.DataReturnMiddleware(&l))
	handler.Use(middlewares.EventsMiddleware(&l, config.GetMustLogInfo(), config.GetMustLogValidationError(), config.GetMustLogServerError()))
	for _, module := range modules {
		module.SetRoutes(handler, &l)
	}

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
