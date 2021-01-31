// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

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

// ChainClient is the client API for Chain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChainClient interface {
	// Node -> Master - Requesting to join the chain
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type chainClient struct {
	cc grpc.ClientConnInterface
}

func NewChainClient(cc grpc.ClientConnInterface) ChainClient {
	return &chainClient{cc}
}

func (c *chainClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/Chain/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChainServer is the server API for Chain service.
// All implementations must embed UnimplementedChainServer
// for forward compatibility
type ChainServer interface {
	// Node -> Master - Requesting to join the chain
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	mustEmbedUnimplementedChainServer()
}

// UnimplementedChainServer must be embedded to have forward compatible implementations.
type UnimplementedChainServer struct {
}

func (UnimplementedChainServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedChainServer) mustEmbedUnimplementedChainServer() {}

// UnsafeChainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChainServer will
// result in compilation errors.
type UnsafeChainServer interface {
	mustEmbedUnimplementedChainServer()
}

func RegisterChainServer(s grpc.ServiceRegistrar, srv ChainServer) {
	s.RegisterService(&Chain_ServiceDesc, srv)
}

func _Chain_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Chain/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Chain_ServiceDesc is the grpc.ServiceDesc for Chain service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chain_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Chain",
	HandlerType: (*ChainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Chain_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/chain.proto",
}
