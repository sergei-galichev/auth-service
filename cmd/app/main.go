package main

import (
	"auth-service/internal/app"
	"auth-service/internal/config"
	"auth-service/pkg/logging"
	"context"
	"os"
	"os/signal"
	"syscall"
)

var (
	ctx context.Context
)

func init() {
	// Initialize empty context
	ctx = context.Background()

	// Initialize configuration
	config.InitDefaultConfig()

	// Initialize logger
	logging.Init()
}

func main() {
	logger := logging.GetLogger()

	a, err := app.NewApp(ctx)
	if err != nil {
		logger.Fatal("Failed to initialize app: ", err)
	}

	err = a.Run()
	if err != nil {
		logger.Fatal("Failed to run app: ", err)
	}

	// Shutdown
	shutdown([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM})
	a.Stop()
}

func shutdown(signals []os.Signal) {
	l := logging.GetLogger()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signals...)
	sig := <-ch
	l.Infof("Caught signal: %s. Shutting down...", sig)
}
