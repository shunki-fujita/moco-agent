// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// HealthServiceClient is the client API for HealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthServiceClient interface {
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type healthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthServiceClient(cc grpc.ClientConnInterface) HealthServiceClient {
	return &healthServiceClient{cc}
}

func (c *healthServiceClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/HealthService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServiceServer is the server API for HealthService service.
// All implementations must embed UnimplementedHealthServiceServer
// for forward compatibility
type HealthServiceServer interface {
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	mustEmbedUnimplementedHealthServiceServer()
}

// UnimplementedHealthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHealthServiceServer struct {
}

func (UnimplementedHealthServiceServer) Health(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedHealthServiceServer) mustEmbedUnimplementedHealthServiceServer() {}

// UnsafeHealthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServiceServer will
// result in compilation errors.
type UnsafeHealthServiceServer interface {
	mustEmbedUnimplementedHealthServiceServer()
}

func RegisterHealthServiceServer(s grpc.ServiceRegistrar, srv HealthServiceServer) {
	s.RegisterService(&HealthService_ServiceDesc, srv)
}

func _HealthService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HealthService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthService_ServiceDesc is the grpc.ServiceDesc for HealthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HealthService",
	HandlerType: (*HealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _HealthService_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}

// CloneServiceClient is the client API for CloneService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloneServiceClient interface {
	Clone(ctx context.Context, in *CloneRequest, opts ...grpc.CallOption) (*CloneResponse, error)
}

type cloneServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCloneServiceClient(cc grpc.ClientConnInterface) CloneServiceClient {
	return &cloneServiceClient{cc}
}

func (c *cloneServiceClient) Clone(ctx context.Context, in *CloneRequest, opts ...grpc.CallOption) (*CloneResponse, error) {
	out := new(CloneResponse)
	err := c.cc.Invoke(ctx, "/CloneService/Clone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloneServiceServer is the server API for CloneService service.
// All implementations must embed UnimplementedCloneServiceServer
// for forward compatibility
type CloneServiceServer interface {
	Clone(context.Context, *CloneRequest) (*CloneResponse, error)
	mustEmbedUnimplementedCloneServiceServer()
}

// UnimplementedCloneServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCloneServiceServer struct {
}

func (UnimplementedCloneServiceServer) Clone(context.Context, *CloneRequest) (*CloneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clone not implemented")
}
func (UnimplementedCloneServiceServer) mustEmbedUnimplementedCloneServiceServer() {}

// UnsafeCloneServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloneServiceServer will
// result in compilation errors.
type UnsafeCloneServiceServer interface {
	mustEmbedUnimplementedCloneServiceServer()
}

func RegisterCloneServiceServer(s grpc.ServiceRegistrar, srv CloneServiceServer) {
	s.RegisterService(&CloneService_ServiceDesc, srv)
}

func _CloneService_Clone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloneServiceServer).Clone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CloneService/Clone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloneServiceServer).Clone(ctx, req.(*CloneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CloneService_ServiceDesc is the grpc.ServiceDesc for CloneService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloneService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CloneService",
	HandlerType: (*CloneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Clone",
			Handler:    _CloneService_Clone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
