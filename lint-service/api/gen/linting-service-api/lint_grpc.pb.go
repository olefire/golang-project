// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: proto/lint.proto

package linting_service_api

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

const (
	LintingService_LintCode_FullMethodName = "/gen.LintingService/LintCode"
)

// LintingServiceClient is the client API for LintingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LintingServiceClient interface {
	LintCode(ctx context.Context, in *LintCodeRequest, opts ...grpc.CallOption) (*LintCodeResponse, error)
}

type lintingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLintingServiceClient(cc grpc.ClientConnInterface) LintingServiceClient {
	return &lintingServiceClient{cc}
}

func (c *lintingServiceClient) LintCode(ctx context.Context, in *LintCodeRequest, opts ...grpc.CallOption) (*LintCodeResponse, error) {
	out := new(LintCodeResponse)
	err := c.cc.Invoke(ctx, LintingService_LintCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LintingServiceServer is the server API for LintingService service.
// All implementations must embed UnimplementedLintingServiceServer
// for forward compatibility
type LintingServiceServer interface {
	LintCode(context.Context, *LintCodeRequest) (*LintCodeResponse, error)
	mustEmbedUnimplementedLintingServiceServer()
}

// UnimplementedLintingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLintingServiceServer struct {
}

func (UnimplementedLintingServiceServer) LintCode(context.Context, *LintCodeRequest) (*LintCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LintCode not implemented")
}
func (UnimplementedLintingServiceServer) mustEmbedUnimplementedLintingServiceServer() {}

// UnsafeLintingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LintingServiceServer will
// result in compilation errors.
type UnsafeLintingServiceServer interface {
	mustEmbedUnimplementedLintingServiceServer()
}

func RegisterLintingServiceServer(s grpc.ServiceRegistrar, srv LintingServiceServer) {
	s.RegisterService(&LintingService_ServiceDesc, srv)
}

func _LintingService_LintCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LintCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LintingServiceServer).LintCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LintingService_LintCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LintingServiceServer).LintCode(ctx, req.(*LintCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LintingService_ServiceDesc is the grpc.ServiceDesc for LintingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LintingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gen.LintingService",
	HandlerType: (*LintingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LintCode",
			Handler:    _LintingService_LintCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/lint.proto",
}
