// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ShittyChatClient is the client API for ShittyChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShittyChatClient interface {
	Publish(ctx context.Context, opts ...grpc.CallOption) (ShittyChat_PublishClient, error)
	Broadcast(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (ShittyChat_BroadcastClient, error)
}

type shittyChatClient struct {
	cc grpc.ClientConnInterface
}

func NewShittyChatClient(cc grpc.ClientConnInterface) ShittyChatClient {
	return &shittyChatClient{cc}
}

func (c *shittyChatClient) Publish(ctx context.Context, opts ...grpc.CallOption) (ShittyChat_PublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &ShittyChat_ServiceDesc.Streams[0], "/service.ShittyChat/Publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &shittyChatPublishClient{stream}
	return x, nil
}

type ShittyChat_PublishClient interface {
	Send(*UserMessage) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type shittyChatPublishClient struct {
	grpc.ClientStream
}

func (x *shittyChatPublishClient) Send(m *UserMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *shittyChatPublishClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *shittyChatClient) Broadcast(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (ShittyChat_BroadcastClient, error) {
	stream, err := c.cc.NewStream(ctx, &ShittyChat_ServiceDesc.Streams[1], "/service.ShittyChat/Broadcast", opts...)
	if err != nil {
		return nil, err
	}
	x := &shittyChatBroadcastClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ShittyChat_BroadcastClient interface {
	Recv() (*UserMessage, error)
	grpc.ClientStream
}

type shittyChatBroadcastClient struct {
	grpc.ClientStream
}

func (x *shittyChatBroadcastClient) Recv() (*UserMessage, error) {
	m := new(UserMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ShittyChatServer is the server API for ShittyChat service.
// All implementations must embed UnimplementedShittyChatServer
// for forward compatibility
type ShittyChatServer interface {
	Publish(ShittyChat_PublishServer) error
	Broadcast(*emptypb.Empty, ShittyChat_BroadcastServer) error
	mustEmbedUnimplementedShittyChatServer()
}

// UnimplementedShittyChatServer must be embedded to have forward compatible implementations.
type UnimplementedShittyChatServer struct {
}

func (UnimplementedShittyChatServer) Publish(ShittyChat_PublishServer) error {
	return status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedShittyChatServer) Broadcast(*emptypb.Empty, ShittyChat_BroadcastServer) error {
	return status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedShittyChatServer) mustEmbedUnimplementedShittyChatServer() {}

// UnsafeShittyChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShittyChatServer will
// result in compilation errors.
type UnsafeShittyChatServer interface {
	mustEmbedUnimplementedShittyChatServer()
}

func RegisterShittyChatServer(s grpc.ServiceRegistrar, srv ShittyChatServer) {
	s.RegisterService(&ShittyChat_ServiceDesc, srv)
}

func _ShittyChat_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ShittyChatServer).Publish(&shittyChatPublishServer{stream})
}

type ShittyChat_PublishServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*UserMessage, error)
	grpc.ServerStream
}

type shittyChatPublishServer struct {
	grpc.ServerStream
}

func (x *shittyChatPublishServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *shittyChatPublishServer) Recv() (*UserMessage, error) {
	m := new(UserMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ShittyChat_Broadcast_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ShittyChatServer).Broadcast(m, &shittyChatBroadcastServer{stream})
}

type ShittyChat_BroadcastServer interface {
	Send(*UserMessage) error
	grpc.ServerStream
}

type shittyChatBroadcastServer struct {
	grpc.ServerStream
}

func (x *shittyChatBroadcastServer) Send(m *UserMessage) error {
	return x.ServerStream.SendMsg(m)
}

// ShittyChat_ServiceDesc is the grpc.ServiceDesc for ShittyChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShittyChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.ShittyChat",
	HandlerType: (*ShittyChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Publish",
			Handler:       _ShittyChat_Publish_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Broadcast",
			Handler:       _ShittyChat_Broadcast_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service/services.proto",
}
