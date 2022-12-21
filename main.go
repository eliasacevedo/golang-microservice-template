package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/eliasacevedo/golang-microservice-template/src/config"
	"github.com/eliasacevedo/golang-microservice-template/src/server"
	"github.com/eliasacevedo/golang-microservice-template/src/utilities"
	"github.com/gin-gonic/gin"
)

func main() {
	// Base logger
	l := utilities.NewLogger()

	// Loading environment variables from external file
	config.LoadEnvFromFile(l)

	// Running server
	gin.SetMode(config.GetAppMode())
	srv := server.NewServer(server.GetServerConfigFromEnvVar())
	go server.RunServer(srv, l)

	// Detect when server will shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	sig := <-c
	server.OnShutDownServer(srv, l, sig)
}
