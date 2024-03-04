// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: jwt.proto

package jwtGRPC

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// JWTServiceClient is the client API for JWTService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JWTServiceClient interface {
	GenerateToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	RefreshToken(ctx context.Context, in *TokenRequestRefresh, opts ...grpc.CallOption) (*TokenResponse, error)
}

type jWTServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJWTServiceClient(cc grpc.ClientConnInterface) JWTServiceClient {
	return &jWTServiceClient{cc}
}

func (c *jWTServiceClient) GenerateToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/jwtGRPC.JWTService/GenerateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jWTServiceClient) RefreshToken(ctx context.Context, in *TokenRequestRefresh, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/jwtGRPC.JWTService/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JWTServiceServer is the server API for JWTService service.
// All implementations must embed UnimplementedJWTServiceServer
// for forward compatibility
type JWTServiceServer interface {
	GenerateToken(context.Context, *TokenRequest) (*TokenResponse, error)
	RefreshToken(context.Context, *TokenRequestRefresh) (*TokenResponse, error)
	mustEmbedUnimplementedJWTServiceServer()
}

// UnimplementedJWTServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJWTServiceServer struct {
}

func (UnimplementedJWTServiceServer) GenerateToken(context.Context, *TokenRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (UnimplementedJWTServiceServer) RefreshToken(context.Context, *TokenRequestRefresh) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedJWTServiceServer) mustEmbedUnimplementedJWTServiceServer() {}

// UnsafeJWTServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JWTServiceServer will
// result in compilation errors.
type UnsafeJWTServiceServer interface {
	mustEmbedUnimplementedJWTServiceServer()
}

func RegisterJWTServiceServer(s grpc.ServiceRegistrar, srv JWTServiceServer) {
	s.RegisterService(&JWTService_ServiceDesc, srv)
}

func _JWTService_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JWTServiceServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jwtGRPC.JWTService/GenerateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JWTServiceServer).GenerateToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JWTService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequestRefresh)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JWTServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jwtGRPC.JWTService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JWTServiceServer).RefreshToken(ctx, req.(*TokenRequestRefresh))
	}
	return interceptor(ctx, in, info, handler)
}

// JWTService_ServiceDesc is the grpc.ServiceDesc for JWTService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JWTService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jwtGRPC.JWTService",
	HandlerType: (*JWTServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateToken",
			Handler:    _JWTService_GenerateToken_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _JWTService_RefreshToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jwt.proto",
}
