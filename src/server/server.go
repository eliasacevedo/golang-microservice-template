package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	config "github.com/eliasacevedo/golang-microservice-template/src/config"
	"github.com/eliasacevedo/golang-microservice-template/src/utilities"
)

type ServerConfig struct {
	Port                     string
	ReadTimeout              time.Duration
	ReadHeaderTimeout        time.Duration
	WriteTimeout             time.Duration
	IdleTimeout              time.Duration
	TimeBeforeShutdownServer time.Duration
}

func NewServer(config ServerConfig) *http.Server {
	srv := &http.Server{
		Addr:              config.Port,
		Handler:           NewRouter(),
		ReadTimeout:       config.ReadTimeout,
		ReadHeaderTimeout: config.ReadHeaderTimeout,
		WriteTimeout:      config.WriteTimeout,
		// BaseContext:       modifyContext,
		IdleTimeout: config.IdleTimeout,
	}

	return srv
}

func GetServerConfigFromEnvVar() ServerConfig {
	return ServerConfig{
		Port:              fmt.Sprintf(":%s", config.GetPort()),
		ReadTimeout:       config.GetReadTimeout() * time.Second,
		ReadHeaderTimeout: config.GetReadHeaderTimeout() * time.Second,
		WriteTimeout:      config.GetWriteTimeout() * time.Second,
		IdleTimeout:       config.GetIdleTimeout() * time.Second,
	}
}

func RunServer(srv *http.Server, l utilities.Logger) {
	l.Info(fmt.Sprintf("Initializing server in: %s", srv.Addr))
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		l.PanicApp(fmt.Sprintf("listen: %s\n", err))
	}
}

func OnShutDownServer(srv *http.Server, l utilities.Logger, sig os.Signal) {
	l.Info(fmt.Sprintf("Shutting down Server: %d", sig))

	t := config.GetTimeBeforeShutdownServer()

	ctx, cancel := context.WithTimeout(context.Background(), t*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		l.PanicApp(fmt.Sprintf("Forced shutdown: %s", err))
	}

	<-ctx.Done()

	l.Info("All pending transactions completed")
	l.Info("Server exiting")
}

// func modifyContext(net.Listener) context.Context {
// 	ctx := context.WithValue(context.Background(), )
// 	return ctx
// }
