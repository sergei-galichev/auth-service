package app

import (
	"auth-service/internal/config"
	auth_v1 "auth-service/internal/delivery/grpc_v1"
	repositories "auth-service/internal/repository"
	userRepo "auth-service/internal/repository/user/postgres"
	services "auth-service/internal/service"
	userService "auth-service/internal/service/user"
	"auth-service/pkg/auth_jwt"
	"auth-service/pkg/cache"
	"auth-service/pkg/cache/redis"
	"auth-service/pkg/logging"
	"auth-service/pkg/storage/postgres"
)

type serviceProvider struct {
	userRepo repositories.UserRepository

	userService services.UserService

	authHelper auth_jwt.JWTHelper

	authImplementation *auth_v1.AuthImplementation

	storage *postgres.Storage

	cache cache.Repository

	logger *logging.Logger

	pgCfg    config.PGConfig
	redisCfg config.RedisConfig
	authCfg  config.AuthConfig
}

func newServiceProvider(
	pgCfg config.PGConfig,
	redisCfg config.RedisConfig,
	authCfg config.AuthConfig,
) *serviceProvider {
	logger := logging.GetLogger()

	return &serviceProvider{
		logger:   logger,
		pgCfg:    pgCfg,
		redisCfg: redisCfg,
		authCfg:  authCfg,
	}
}

func (s *serviceProvider) Storage() *postgres.Storage {
	var err error
	if s.storage == nil {
		s.storage, err = postgres.New(s.pgCfg.DSN())
		if err != nil {
			s.logger.Fatal("Failed to connect to postgres: ", err)
		}
	}
	return s.storage
}

func (s *serviceProvider) UserRepository() repositories.UserRepository {
	if s.userRepo == nil {
		s.userRepo = userRepo.New(s.Storage().Session)
	}
	s.logger.Debug("UserRepository created")

	return s.userRepo
}

func (s *serviceProvider) UserCache() cache.Repository {
	if s.cache == nil {
		s.cache = redis.NewCache(s.redisCfg.Address())
	}
	s.logger.Debug("UserCache created")

	return s.cache
}

func (s *serviceProvider) AuthHelper() auth_jwt.JWTHelper {
	if s.authHelper == nil {
		s.authHelper = auth_jwt.NewHelper(s.authCfg)
	}

	return s.authHelper
}

func (s *serviceProvider) UserService() services.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(),
			s.UserCache(),
			s.AuthHelper(),
		)
	}
	s.logger.Debug("UserService created")

	return s.userService
}

func (s *serviceProvider) AuthImplementation() {
	if s.authImplementation == nil {
		s.authImplementation = auth_v1.NewImplementation(
			s.UserService(),
		)
	}
	s.logger.Debug("AuthImplementation created")
}
