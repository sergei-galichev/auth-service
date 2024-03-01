package main

import (
	"auth-service/internal/app"
	"auth-service/internal/config"
	"auth-service/pkg/logging"
	"context"
	"log"
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
	//config.InitDefaultConfig()
	err := config.LoadConfig("dev.env")
	if err != nil {
		log.Println("Couldn't load env file")
	}

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
