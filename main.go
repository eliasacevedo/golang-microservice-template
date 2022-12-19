package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eliasacevedo/golang-microservice-template/src/utilities"
	"github.com/joho/godotenv"
)

func main() {
	environment := os.Getenv("env")
	if environment == "" {
		environment = ".env.local"
	}

	l := utilities.NewLogger()

	err := godotenv.Load(environment)
	if err != nil {
		l.PanicApp(fmt.Sprintf("error loading env file: %s", err.Error()))
	}

	port := os.Getenv("PORT")
	address := os.Getenv("ADDRESS")

	l.Info(fmt.Sprintf("running in %s:%s", address, port))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	sig := <-c
	l.Info(fmt.Sprintf("Got signal: %d", sig))

	context.WithTimeout(context.Background(), 30*time.Second)
}
