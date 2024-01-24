package grpc_v1

import (
	auth_v1 "auth-service/pkg/grpc/v1/auth"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *AuthImplementation) Register(ctx context.Context, req *auth_v1.RegisterRequest) (*empty.Empty, error) {
	_ = ctx
	pass := req.GetPassword()
	confirmPass := req.GetConfirmPassword()
	if pass != confirmPass {
		return &empty.Empty{}, status.Error(codes.InvalidArgument, "Password mismatch")
	}

	email := req.GetEmail()
	if !i.userService.CheckEmail(email) {
		return &empty.Empty{}, status.Error(codes.InvalidArgument, "Email format error")
	}

	//i.userService.CreateUser()
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
