package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	environment := os.Getenv("env")
	if environment == "" {
		environment = ".env.local"
	}

	err := godotenv.Load(environment)
	if err != nil {
		log.Fatal("error loading env file: ", err)
	}

	port := os.Getenv("PORT")
	address := os.Getenv("ADDRESS")

	log.Printf("running in %s:%s", address, port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	sig := <-c
	log.Println("Got signal:", sig)

	context.WithTimeout(context.Background(), 30*time.Second)
}
