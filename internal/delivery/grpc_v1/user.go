package grpc_v1

import (
	"auth-service/internal/delivery/grpc_v1/dto"
	auth_v1 "auth-service/pkg/grpc/v1/auth"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *AuthImplementation) Register(ctx context.Context, req *auth_v1.RegisterRequest) (
	*auth_v1.RegisterResponse,
	error,
) {
	_ = ctx

	id, err := i.userService.CreateUser(
		&dto.UserCreateDTO{
			Email:           req.GetEmail(),
			Password:        req.GetPassword(),
			ConfirmPassword: req.GetConfirmPassword(),
			Role:            req.GetRole().String(),
			AdminKey:        req.GetAdminKey(),
		},
	)
	if err != nil {
		return &auth_v1.RegisterResponse{
			Status:  "failed",
			Message: err.Error(),
		}, status.Error(codes.Internal, err.Error())
	}

	return &auth_v1.RegisterResponse{
		Status:  "success",
		Message: fmt.Sprintf("User registered. ID: %d", id),
	}, nil
}

func (i *AuthImplementation) Login(ctx context.Context, req *auth_v1.LoginRequest) (
	*auth_v1.LoginResponse,
	error,
) {
	_ = ctx

	at, rt, err := i.userService.LoginUser(
		&dto.UserLoginDTO{
			Email:    req.GetEmail(),
			Password: req.GetPassword(),
		},
	)

	if err != nil {
		return &auth_v1.LoginResponse{
			AccessToken:  "",
			RefreshToken: "",
		}, status.Error(codes.Internal, err.Error())
	}

	return &auth_v1.LoginResponse{
		AccessToken:  at,
		RefreshToken: rt,
	}, nil
}

func (i *AuthImplementation) Logout(ctx context.Context, req *auth_v1.LogoutRequest) (*empty.Empty, error) {
	_ = ctx
	err := i.userService.LogoutUser(
		&dto.UserLogoutDTO{
			AccessToken:  req.GetAccessToken(),
			RefreshToken: req.GetRefreshToken(),
		},
	)

	if err != nil {
		return &empty.Empty{}, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}
