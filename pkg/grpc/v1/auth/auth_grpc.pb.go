// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: auth/auth.proto

package auth_v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthServiceV1Client is the client API for AuthServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceV1Client interface {
	// Register used to user registration
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// Login used to user authentication
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// Logout used to user log out
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type authServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceV1Client(cc grpc.ClientConnInterface) AuthServiceV1Client {
	return &authServiceV1Client{cc}
}

func (c *authServiceV1Client) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/auth_v1.AuthServiceV1/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceV1Client) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth_v1.AuthServiceV1/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceV1Client) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/auth_v1.AuthServiceV1/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceV1Server is the server API for AuthServiceV1 service.
// All implementations must embed UnimplementedAuthServiceV1Server
// for forward compatibility
type AuthServiceV1Server interface {
	// Register used to user registration
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// Login used to user authentication
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Logout used to user log out
	Logout(context.Context, *LogoutRequest) (*empty.Empty, error)
	mustEmbedUnimplementedAuthServiceV1Server()
}

// UnimplementedAuthServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceV1Server struct {
}

func (UnimplementedAuthServiceV1Server) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServiceV1Server) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceV1Server) Logout(context.Context, *LogoutRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedAuthServiceV1Server) mustEmbedUnimplementedAuthServiceV1Server() {}

// UnsafeAuthServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceV1Server will
// result in compilation errors.
type UnsafeAuthServiceV1Server interface {
	mustEmbedUnimplementedAuthServiceV1Server()
}

func RegisterAuthServiceV1Server(s grpc.ServiceRegistrar, srv AuthServiceV1Server) {
	s.RegisterService(&AuthServiceV1_ServiceDesc, srv)
}

func _AuthServiceV1_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceV1Server).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_v1.AuthServiceV1/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceV1Server).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceV1_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceV1Server).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_v1.AuthServiceV1/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceV1Server).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServiceV1_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceV1Server).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_v1.AuthServiceV1/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceV1Server).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthServiceV1_ServiceDesc is the grpc.ServiceDesc for AuthServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_v1.AuthServiceV1",
	HandlerType: (*AuthServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AuthServiceV1_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthServiceV1_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _AuthServiceV1_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}
