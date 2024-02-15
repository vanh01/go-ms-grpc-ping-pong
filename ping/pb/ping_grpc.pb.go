// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: ping.proto

package pb

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
	Ping_PingMessage_FullMethodName = "/Ping/PingMessage"
)

// PingClient is the client API for Ping service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PingClient interface {
	PingMessage(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type pingClient struct {
	cc grpc.ClientConnInterface
}

func NewPingClient(cc grpc.ClientConnInterface) PingClient {
	return &pingClient{cc}
}

func (c *pingClient) PingMessage(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, Ping_PingMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingServer is the server API for Ping service.
// All implementations must embed UnimplementedPingServer
// for forward compatibility
type PingServer interface {
	PingMessage(context.Context, *PingRequest) (*PingResponse, error)
	mustEmbedUnimplementedPingServer()
}

// UnimplementedPingServer must be embedded to have forward compatible implementations.
type UnimplementedPingServer struct {
}

func (UnimplementedPingServer) PingMessage(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingMessage not implemented")
}
func (UnimplementedPingServer) mustEmbedUnimplementedPingServer() {}

// UnsafePingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PingServer will
// result in compilation errors.
type UnsafePingServer interface {
	mustEmbedUnimplementedPingServer()
}

func RegisterPingServer(s grpc.ServiceRegistrar, srv PingServer) {
	s.RegisterService(&Ping_ServiceDesc, srv)
}

func _Ping_PingMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServer).PingMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ping_PingMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServer).PingMessage(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Ping_ServiceDesc is the grpc.ServiceDesc for Ping service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ping_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ping",
	HandlerType: (*PingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingMessage",
			Handler:    _Ping_PingMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ping.proto",
}