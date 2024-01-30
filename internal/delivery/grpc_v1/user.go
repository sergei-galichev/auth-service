package grpc_v1

import (
	"auth-service/internal/delivery/grpc_v1/dto"
	auth_v1 "auth-service/pkg/grpc/v1/auth"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func (i *AuthImplementation) Register(ctx context.Context, req *auth_v1.RegisterRequest) (*empty.Empty, error) {
	_ = ctx

	_, err := i.userService.CreateUser(
		&dto.UserCreateDTO{
			Email:           req.GetEmail(),
			Password:        req.GetPassword(),
			ConfirmPassword: req.GetConfirmPassword(),
			Role:            req.GetRole().String(),
			AdminKey:        req.GetAdminKey(),
		},
	)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func FakeAuthUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	return handler(ctx, req)
}
