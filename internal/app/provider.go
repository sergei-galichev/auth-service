package app

import (
	auth_v1 "auth-service/internal/delivery/grpc_v1"
	repositories "auth-service/internal/repository"
	userRepo "auth-service/internal/repository/user/postgres"
	services "auth-service/internal/service"
	userService "auth-service/internal/service/user"
	"auth-service/pkg/cache"
	"auth-service/pkg/cache/redis"
	"auth-service/pkg/logging"
	"auth-service/pkg/storage/postgres"
)

type serviceProvider struct {
	userRepo repositories.UserRepository

	userService services.UserService

	authImplementation *auth_v1.AuthImplementation

	storage *postgres.Storage

	cache cache.Repository

	logger *logging.Logger
}

func newServiceProvider(dsn string) *serviceProvider {
	logger := logging.GetLogger()
	storage, err := postgres.New(dsn)
	if err != nil {
		logger.Fatal("Failed to connect to postgres: ", err)
	}
	return &serviceProvider{
		storage: storage,
		logger:  logger,
	}
}

func (s *serviceProvider) UserRepository() repositories.UserRepository {
	if s.userRepo == nil {
		s.userRepo = userRepo.New(s.storage.Pool)
	}
	s.logger.Debug("UserRepository created")

	return s.userRepo
}

func (s *serviceProvider) UserCache() cache.Repository {
	if s.cache == nil {
		s.cache = redis.NewCache()
	}
	s.logger.Debug("UserCache created")

	return s.cache
}

func (s *serviceProvider) UserService() services.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(),
			s.UserCache(),
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
