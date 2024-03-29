// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: thesocialrobot/thesocialrobot.proto

package thesocialrobot

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

// TheSocialRobotClient is the client API for TheSocialRobot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TheSocialRobotClient interface {
	EventStream(ctx context.Context, opts ...grpc.CallOption) (TheSocialRobot_EventStreamClient, error)
}

type theSocialRobotClient struct {
	cc grpc.ClientConnInterface
}

func NewTheSocialRobotClient(cc grpc.ClientConnInterface) TheSocialRobotClient {
	return &theSocialRobotClient{cc}
}

func (c *theSocialRobotClient) EventStream(ctx context.Context, opts ...grpc.CallOption) (TheSocialRobot_EventStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &TheSocialRobot_ServiceDesc.Streams[0], "/thesocialrobot.TheSocialRobot/EventStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &theSocialRobotEventStreamClient{stream}
	return x, nil
}

type TheSocialRobot_EventStreamClient interface {
	Send(*BodyEvent) error
	Recv() (*BodyCommand, error)
	grpc.ClientStream
}

type theSocialRobotEventStreamClient struct {
	grpc.ClientStream
}

func (x *theSocialRobotEventStreamClient) Send(m *BodyEvent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *theSocialRobotEventStreamClient) Recv() (*BodyCommand, error) {
	m := new(BodyCommand)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TheSocialRobotServer is the server API for TheSocialRobot service.
// All implementations must embed UnimplementedTheSocialRobotServer
// for forward compatibility
type TheSocialRobotServer interface {
	EventStream(TheSocialRobot_EventStreamServer) error
	mustEmbedUnimplementedTheSocialRobotServer()
}

// UnimplementedTheSocialRobotServer must be embedded to have forward compatible implementations.
type UnimplementedTheSocialRobotServer struct {
}

func (UnimplementedTheSocialRobotServer) EventStream(TheSocialRobot_EventStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EventStream not implemented")
}
func (UnimplementedTheSocialRobotServer) mustEmbedUnimplementedTheSocialRobotServer() {}

// UnsafeTheSocialRobotServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TheSocialRobotServer will
// result in compilation errors.
type UnsafeTheSocialRobotServer interface {
	mustEmbedUnimplementedTheSocialRobotServer()
}

func RegisterTheSocialRobotServer(s grpc.ServiceRegistrar, srv TheSocialRobotServer) {
	s.RegisterService(&TheSocialRobot_ServiceDesc, srv)
}

func _TheSocialRobot_EventStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TheSocialRobotServer).EventStream(&theSocialRobotEventStreamServer{stream})
}

type TheSocialRobot_EventStreamServer interface {
	Send(*BodyCommand) error
	Recv() (*BodyEvent, error)
	grpc.ServerStream
}

type theSocialRobotEventStreamServer struct {
	grpc.ServerStream
}

func (x *theSocialRobotEventStreamServer) Send(m *BodyCommand) error {
	return x.ServerStream.SendMsg(m)
}

func (x *theSocialRobotEventStreamServer) Recv() (*BodyEvent, error) {
	m := new(BodyEvent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TheSocialRobot_ServiceDesc is the grpc.ServiceDesc for TheSocialRobot service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TheSocialRobot_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "thesocialrobot.TheSocialRobot",
	HandlerType: (*TheSocialRobotServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EventStream",
			Handler:       _TheSocialRobot_EventStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "thesocialrobot/thesocialrobot.proto",
}
