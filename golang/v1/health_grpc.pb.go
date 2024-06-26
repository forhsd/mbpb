// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: health.proto

package mbpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	MBLink_Health_FullMethodName = "/mbpb.MBLink/Health"
	MBLink_ReView_FullMethodName = "/mbpb.MBLink/ReView"
)

// MBLinkClient is the client API for MBLink service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MBLinkClient interface {
	// 健康检查
	Health(ctx context.Context, opts ...grpc.CallOption) (MBLink_HealthClient, error)
	// 运行情况
	ReView(ctx context.Context, opts ...grpc.CallOption) (MBLink_ReViewClient, error)
}

type mBLinkClient struct {
	cc grpc.ClientConnInterface
}

func NewMBLinkClient(cc grpc.ClientConnInterface) MBLinkClient {
	return &mBLinkClient{cc}
}

func (c *mBLinkClient) Health(ctx context.Context, opts ...grpc.CallOption) (MBLink_HealthClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MBLink_ServiceDesc.Streams[0], MBLink_Health_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &mBLinkHealthClient{ClientStream: stream}
	return x, nil
}

type MBLink_HealthClient interface {
	Send(*HealthRequest) error
	Recv() (*HealthReply, error)
	grpc.ClientStream
}

type mBLinkHealthClient struct {
	grpc.ClientStream
}

func (x *mBLinkHealthClient) Send(m *HealthRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mBLinkHealthClient) Recv() (*HealthReply, error) {
	m := new(HealthReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mBLinkClient) ReView(ctx context.Context, opts ...grpc.CallOption) (MBLink_ReViewClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MBLink_ServiceDesc.Streams[1], MBLink_ReView_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &mBLinkReViewClient{ClientStream: stream}
	return x, nil
}

type MBLink_ReViewClient interface {
	Send(*Overview) error
	Recv() (*Ack, error)
	grpc.ClientStream
}

type mBLinkReViewClient struct {
	grpc.ClientStream
}

func (x *mBLinkReViewClient) Send(m *Overview) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mBLinkReViewClient) Recv() (*Ack, error) {
	m := new(Ack)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MBLinkServer is the server API for MBLink service.
// All implementations must embed UnimplementedMBLinkServer
// for forward compatibility
type MBLinkServer interface {
	// 健康检查
	Health(MBLink_HealthServer) error
	// 运行情况
	ReView(MBLink_ReViewServer) error
	mustEmbedUnimplementedMBLinkServer()
}

// UnimplementedMBLinkServer must be embedded to have forward compatible implementations.
type UnimplementedMBLinkServer struct {
}

func (UnimplementedMBLinkServer) Health(MBLink_HealthServer) error {
	return status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedMBLinkServer) ReView(MBLink_ReViewServer) error {
	return status.Errorf(codes.Unimplemented, "method ReView not implemented")
}
func (UnimplementedMBLinkServer) mustEmbedUnimplementedMBLinkServer() {}

// UnsafeMBLinkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MBLinkServer will
// result in compilation errors.
type UnsafeMBLinkServer interface {
	mustEmbedUnimplementedMBLinkServer()
}

func RegisterMBLinkServer(s grpc.ServiceRegistrar, srv MBLinkServer) {
	s.RegisterService(&MBLink_ServiceDesc, srv)
}

func _MBLink_Health_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MBLinkServer).Health(&mBLinkHealthServer{ServerStream: stream})
}

type MBLink_HealthServer interface {
	Send(*HealthReply) error
	Recv() (*HealthRequest, error)
	grpc.ServerStream
}

type mBLinkHealthServer struct {
	grpc.ServerStream
}

func (x *mBLinkHealthServer) Send(m *HealthReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mBLinkHealthServer) Recv() (*HealthRequest, error) {
	m := new(HealthRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MBLink_ReView_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MBLinkServer).ReView(&mBLinkReViewServer{ServerStream: stream})
}

type MBLink_ReViewServer interface {
	Send(*Ack) error
	Recv() (*Overview, error)
	grpc.ServerStream
}

type mBLinkReViewServer struct {
	grpc.ServerStream
}

func (x *mBLinkReViewServer) Send(m *Ack) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mBLinkReViewServer) Recv() (*Overview, error) {
	m := new(Overview)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MBLink_ServiceDesc is the grpc.ServiceDesc for MBLink service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MBLink_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mbpb.MBLink",
	HandlerType: (*MBLinkServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Health",
			Handler:       _MBLink_Health_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ReView",
			Handler:       _MBLink_ReView_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "health.proto",
}
