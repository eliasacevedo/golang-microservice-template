package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/eliasacevedo/golang-microservice-template/src/config"
	"github.com/eliasacevedo/golang-microservice-template/src/core"
	"github.com/eliasacevedo/golang-microservice-template/src/server"
	"github.com/eliasacevedo/golang-microservice-template/src/utilities"
	e "github.com/eliasacevedo/golang-microservice-template/src/x/module"
)

func main() {
	// Base logger
	l := utilities.NewLogger()

	// Loading environment variables from external file
	config.LoadEnvFromFile(l)

	modules := CreateAllModules()
	handler := server.NewRouter(&l)
	for _, module := range modules {
		module.SetRoutes(handler, &l)
		// add middlewares
		// add services
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

func CreateAllModules() []core.Module {
	ms := []core.Module{
		// Add here all modules implementation
		e.NewModule(),
	}
	return ms
}
