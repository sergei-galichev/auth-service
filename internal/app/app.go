package app

import (
	"auth-service/internal/config"
	"auth-service/internal/config/env"
	auth_v1 "auth-service/pkg/grpc/v1/auth"
	"auth-service/pkg/logging"
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type App struct {
	grpcServer *grpc.Server
	options    []grpc.ServerOption

	pgCfg    config.PGConfig
	redisCfg config.RedisConfig
	grpcCfg  config.GRPCConfig
	authCfg  config.AuthConfig

	provider *serviceProvider

	logger *logging.Logger
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{
		logger: logging.GetLogger(),
	}

	err := app.initDeps(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return app, nil
}

func (app *App) Run() error {
	err := app.runGRPCServer()
	if err != nil {
		app.logger.Error("Failed to run gRPC server: ", err)
		return errors.WithStack(err)
	}
	return nil
}

func (app *App) Stop() {
	app.provider.storage.Close()
	app.stopGRPCServer()
}

func (app *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		app.initConfig,
		app.initServiceProvider,
		app.initGRPCServer,
	}

	for _, fn := range inits {
		if err := fn(ctx); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (app *App) initConfig(_ context.Context) error {
	pgCfg, err := env.NewPGConfig()
	if err != nil {
		app.logger.Error("Failed to load postgres config: ", err)
		return errors.WithStack(err)
	}

	redisCfg, err := env.NewRedisConfig()
	if err != nil {
		app.logger.Error("Failed to load redis config: ", err)
		return errors.WithStack(err)
	}

	grpcCfg, err := env.NewGRPCConfig()
	if err != nil {
		app.logger.Error("Failed to load grpc config: ", err)
		return errors.WithStack(err)
	}

	authCfg, err := env.NewAuthConfig()
	if err != nil {
		app.logger.Error("Failed to load auth config: ", err)
		return errors.WithStack(err)
	}

	app.pgCfg = pgCfg
	app.grpcCfg = grpcCfg
	app.authCfg = authCfg
	app.redisCfg = redisCfg

	return nil
}

func (app *App) initServiceProvider(_ context.Context) error {
	app.provider = newServiceProvider(app.pgCfg, app.redisCfg, app.authCfg)
	app.logger.Debug("ServiceProvider initialized")
	app.provider.AuthImplementation()
	app.logger.Debug("AuthImplementation initialized")

	return nil
}

// initGRPCServer initializes gRPC server
func (app *App) initGRPCServer(_ context.Context) error {
	app.grpcServer = grpc.NewServer(app.options...)
	app.logger.Debug("gRPC Server initialized")

	reflection.Register(app.grpcServer)
	app.logger.Debug("gRPC Server registered reflection")

	auth_v1.RegisterAuthServiceV1Server(app.grpcServer, app.provider.authImplementation)
	app.logger.Debug("gRPC Server and AuthImplementation registered")

	return nil
}

func (app *App) runGRPCServer() error {
	lis, err := net.Listen("tcp", app.grpcCfg.Address())
	if err != nil {
		app.logger.Error("Failed to listen: ", err)
		return errors.WithStack(err)
	}
	go func() {
		app.logger.Info("gRPC Server is running: ", app.grpcCfg.Address())
		if serveErr := app.grpcServer.Serve(lis); serveErr != nil {
			app.logger.Fatal("Failed to serve: ", serveErr)
		}
	}()
	return nil
}

func (app *App) stopGRPCServer() {
	app.logger.Info("gRPC Server is stopping")
	app.grpcServer.GracefulStop()
}
