// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: entrypoint.proto

package mbpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MBetl_Enable_FullMethodName            = "/mbpb.MBetl/Enable"
	MBetl_Disable_FullMethodName           = "/mbpb.MBetl/Disable"
	MBetl_Run_FullMethodName               = "/mbpb.MBetl/Run"
	MBetl_Cancel_FullMethodName            = "/mbpb.MBetl/Cancel"
	MBetl_Remove_FullMethodName            = "/mbpb.MBetl/Remove"
	MBetl_DataLineage_FullMethodName       = "/mbpb.MBetl/DataLineage"
	MBetl_TaskflowEnable_FullMethodName    = "/mbpb.MBetl/TaskflowEnable"
	MBetl_GetTaskflowSpec_FullMethodName   = "/mbpb.MBetl/GetTaskflowSpec"
	MBetl_GetTaskflowStatus_FullMethodName = "/mbpb.MBetl/GetTaskflowStatus"
)

// MBetlClient is the client API for MBetl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MBetlClient interface {
	// 启用
	Enable(ctx context.Context, in *EnableRequest, opts ...grpc.CallOption) (*EnableReply, error)
	// 禁用
	Disable(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	// 运行
	Run(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	// 取消
	Cancel(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	// 删除
	Remove(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Error, error)
	// 数据血亲
	DataLineage(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Graph, error)
	// 任务流启用
	TaskflowEnable(ctx context.Context, in *TaskflowRequest, opts ...grpc.CallOption) (*Error, error)
	// 查询工作流定义
	GetTaskflowSpec(ctx context.Context, in *FlowRequest, opts ...grpc.CallOption) (*Graph, error)
	// 查询工作流状态
	GetTaskflowStatus(ctx context.Context, in *FlowRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Graph], error)
}

type mBetlClient struct {
	cc grpc.ClientConnInterface
}

func NewMBetlClient(cc grpc.ClientConnInterface) MBetlClient {
	return &mBetlClient{cc}
}

func (c *mBetlClient) Enable(ctx context.Context, in *EnableRequest, opts ...grpc.CallOption) (*EnableReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EnableReply)
	err := c.cc.Invoke(ctx, MBetl_Enable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mBetlClient) Disable(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Reply)
	err := c.cc.Invoke(ctx, MBetl_Disable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mBetlClient) Run(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Reply)
	err := c.cc.Invoke(ctx, MBetl_Run_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mBetlClient) Cancel(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Reply)
	err := c.cc.Invoke(ctx, MBetl_Cancel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mBetlClient) Remove(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Error, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Error)
	err := c.cc.Invoke(ctx, MBetl_Remove_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mBetlClient) DataLineage(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Graph, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Graph)
	err := c.cc.Invoke(ctx, MBetl_DataLineage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mBetlClient) TaskflowEnable(ctx context.Context, in *TaskflowRequest, opts ...grpc.CallOption) (*Error, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Error)
	err := c.cc.Invoke(ctx, MBetl_TaskflowEnable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mBetlClient) GetTaskflowSpec(ctx context.Context, in *FlowRequest, opts ...grpc.CallOption) (*Graph, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Graph)
	err := c.cc.Invoke(ctx, MBetl_GetTaskflowSpec_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mBetlClient) GetTaskflowStatus(ctx context.Context, in *FlowRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Graph], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MBetl_ServiceDesc.Streams[0], MBetl_GetTaskflowStatus_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[FlowRequest, Graph]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MBetl_GetTaskflowStatusClient = grpc.ServerStreamingClient[Graph]

// MBetlServer is the server API for MBetl service.
// All implementations must embed UnimplementedMBetlServer
// for forward compatibility.
type MBetlServer interface {
	// 启用
	Enable(context.Context, *EnableRequest) (*EnableReply, error)
	// 禁用
	Disable(context.Context, *Request) (*Reply, error)
	// 运行
	Run(context.Context, *Request) (*Reply, error)
	// 取消
	Cancel(context.Context, *Request) (*Reply, error)
	// 删除
	Remove(context.Context, *Request) (*Error, error)
	// 数据血亲
	DataLineage(context.Context, *Request) (*Graph, error)
	// 任务流启用
	TaskflowEnable(context.Context, *TaskflowRequest) (*Error, error)
	// 查询工作流定义
	GetTaskflowSpec(context.Context, *FlowRequest) (*Graph, error)
	// 查询工作流状态
	GetTaskflowStatus(*FlowRequest, grpc.ServerStreamingServer[Graph]) error
	mustEmbedUnimplementedMBetlServer()
}

// UnimplementedMBetlServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMBetlServer struct{}

func (UnimplementedMBetlServer) Enable(context.Context, *EnableRequest) (*EnableReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedMBetlServer) Disable(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedMBetlServer) Run(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedMBetlServer) Cancel(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (UnimplementedMBetlServer) Remove(context.Context, *Request) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedMBetlServer) DataLineage(context.Context, *Request) (*Graph, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataLineage not implemented")
}
func (UnimplementedMBetlServer) TaskflowEnable(context.Context, *TaskflowRequest) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskflowEnable not implemented")
}
func (UnimplementedMBetlServer) GetTaskflowSpec(context.Context, *FlowRequest) (*Graph, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskflowSpec not implemented")
}
func (UnimplementedMBetlServer) GetTaskflowStatus(*FlowRequest, grpc.ServerStreamingServer[Graph]) error {
	return status.Errorf(codes.Unimplemented, "method GetTaskflowStatus not implemented")
}
func (UnimplementedMBetlServer) mustEmbedUnimplementedMBetlServer() {}
func (UnimplementedMBetlServer) testEmbeddedByValue()               {}

// UnsafeMBetlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MBetlServer will
// result in compilation errors.
type UnsafeMBetlServer interface {
	mustEmbedUnimplementedMBetlServer()
}

func RegisterMBetlServer(s grpc.ServiceRegistrar, srv MBetlServer) {
	// If the following call pancis, it indicates UnimplementedMBetlServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MBetl_ServiceDesc, srv)
}

func _MBetl_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MBetlServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MBetl_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MBetlServer).Enable(ctx, req.(*EnableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MBetl_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MBetlServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MBetl_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MBetlServer).Disable(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MBetl_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MBetlServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MBetl_Run_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MBetlServer).Run(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MBetl_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MBetlServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MBetl_Cancel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MBetlServer).Cancel(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MBetl_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MBetlServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MBetl_Remove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MBetlServer).Remove(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MBetl_DataLineage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MBetlServer).DataLineage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MBetl_DataLineage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MBetlServer).DataLineage(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _MBetl_TaskflowEnable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MBetlServer).TaskflowEnable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MBetl_TaskflowEnable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MBetlServer).TaskflowEnable(ctx, req.(*TaskflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MBetl_GetTaskflowSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MBetlServer).GetTaskflowSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MBetl_GetTaskflowSpec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MBetlServer).GetTaskflowSpec(ctx, req.(*FlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MBetl_GetTaskflowStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FlowRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MBetlServer).GetTaskflowStatus(m, &grpc.GenericServerStream[FlowRequest, Graph]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MBetl_GetTaskflowStatusServer = grpc.ServerStreamingServer[Graph]

// MBetl_ServiceDesc is the grpc.ServiceDesc for MBetl service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MBetl_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mbpb.MBetl",
	HandlerType: (*MBetlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enable",
			Handler:    _MBetl_Enable_Handler,
		},
		{
			MethodName: "Disable",
			Handler:    _MBetl_Disable_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _MBetl_Run_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _MBetl_Cancel_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _MBetl_Remove_Handler,
		},
		{
			MethodName: "DataLineage",
			Handler:    _MBetl_DataLineage_Handler,
		},
		{
			MethodName: "TaskflowEnable",
			Handler:    _MBetl_TaskflowEnable_Handler,
		},
		{
			MethodName: "GetTaskflowSpec",
			Handler:    _MBetl_GetTaskflowSpec_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTaskflowStatus",
			Handler:       _MBetl_GetTaskflowStatus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "entrypoint.proto",
}
