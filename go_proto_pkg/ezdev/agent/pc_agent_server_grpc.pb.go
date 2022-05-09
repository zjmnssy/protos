// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.8.0
// source: ezdev/agent/pc_agent_server.proto

package agent

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

// PcAgentClient is the client API for PcAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PcAgentClient interface {
	GetContainerList(ctx context.Context, in *GetContainerListRequest, opts ...grpc.CallOption) (*GetContainerListResponse, error)
	CreateContainer(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	StartContainer(ctx context.Context, in *StartContainerRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	RestartContainer(ctx context.Context, in *RestartContainerRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	StopContainer(ctx context.Context, in *StopContainerRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	DeleteContainer(ctx context.Context, in *DeleteContainerRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	GetStatus(ctx context.Context, opts ...grpc.CallOption) (PcAgent_GetStatusClient, error)
}

type pcAgentClient struct {
	cc grpc.ClientConnInterface
}

func NewPcAgentClient(cc grpc.ClientConnInterface) PcAgentClient {
	return &pcAgentClient{cc}
}

func (c *pcAgentClient) GetContainerList(ctx context.Context, in *GetContainerListRequest, opts ...grpc.CallOption) (*GetContainerListResponse, error) {
	out := new(GetContainerListResponse)
	err := c.cc.Invoke(ctx, "/agent.PcAgent/GetContainerList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pcAgentClient) CreateContainer(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/agent.PcAgent/CreateContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pcAgentClient) StartContainer(ctx context.Context, in *StartContainerRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/agent.PcAgent/StartContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pcAgentClient) RestartContainer(ctx context.Context, in *RestartContainerRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/agent.PcAgent/RestartContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pcAgentClient) StopContainer(ctx context.Context, in *StopContainerRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/agent.PcAgent/StopContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pcAgentClient) DeleteContainer(ctx context.Context, in *DeleteContainerRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/agent.PcAgent/DeleteContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pcAgentClient) GetStatus(ctx context.Context, opts ...grpc.CallOption) (PcAgent_GetStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &PcAgent_ServiceDesc.Streams[0], "/agent.PcAgent/GetStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &pcAgentGetStatusClient{stream}
	return x, nil
}

type PcAgent_GetStatusClient interface {
	Send(*GetStatusRequest) error
	Recv() (*GetStatsResponse, error)
	grpc.ClientStream
}

type pcAgentGetStatusClient struct {
	grpc.ClientStream
}

func (x *pcAgentGetStatusClient) Send(m *GetStatusRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pcAgentGetStatusClient) Recv() (*GetStatsResponse, error) {
	m := new(GetStatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PcAgentServer is the server API for PcAgent service.
// All implementations must embed UnimplementedPcAgentServer
// for forward compatibility
type PcAgentServer interface {
	GetContainerList(context.Context, *GetContainerListRequest) (*GetContainerListResponse, error)
	CreateContainer(context.Context, *CreateContainerRequest) (*CommonResponse, error)
	StartContainer(context.Context, *StartContainerRequest) (*CommonResponse, error)
	RestartContainer(context.Context, *RestartContainerRequest) (*CommonResponse, error)
	StopContainer(context.Context, *StopContainerRequest) (*CommonResponse, error)
	DeleteContainer(context.Context, *DeleteContainerRequest) (*CommonResponse, error)
	GetStatus(PcAgent_GetStatusServer) error
	mustEmbedUnimplementedPcAgentServer()
}

// UnimplementedPcAgentServer must be embedded to have forward compatible implementations.
type UnimplementedPcAgentServer struct {
}

func (UnimplementedPcAgentServer) GetContainerList(context.Context, *GetContainerListRequest) (*GetContainerListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContainerList not implemented")
}
func (UnimplementedPcAgentServer) CreateContainer(context.Context, *CreateContainerRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContainer not implemented")
}
func (UnimplementedPcAgentServer) StartContainer(context.Context, *StartContainerRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartContainer not implemented")
}
func (UnimplementedPcAgentServer) RestartContainer(context.Context, *RestartContainerRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestartContainer not implemented")
}
func (UnimplementedPcAgentServer) StopContainer(context.Context, *StopContainerRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopContainer not implemented")
}
func (UnimplementedPcAgentServer) DeleteContainer(context.Context, *DeleteContainerRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContainer not implemented")
}
func (UnimplementedPcAgentServer) GetStatus(PcAgent_GetStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedPcAgentServer) mustEmbedUnimplementedPcAgentServer() {}

// UnsafePcAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PcAgentServer will
// result in compilation errors.
type UnsafePcAgentServer interface {
	mustEmbedUnimplementedPcAgentServer()
}

func RegisterPcAgentServer(s grpc.ServiceRegistrar, srv PcAgentServer) {
	s.RegisterService(&PcAgent_ServiceDesc, srv)
}

func _PcAgent_GetContainerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContainerListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PcAgentServer).GetContainerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.PcAgent/GetContainerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PcAgentServer).GetContainerList(ctx, req.(*GetContainerListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PcAgent_CreateContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PcAgentServer).CreateContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.PcAgent/CreateContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PcAgentServer).CreateContainer(ctx, req.(*CreateContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PcAgent_StartContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PcAgentServer).StartContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.PcAgent/StartContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PcAgentServer).StartContainer(ctx, req.(*StartContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PcAgent_RestartContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PcAgentServer).RestartContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.PcAgent/RestartContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PcAgentServer).RestartContainer(ctx, req.(*RestartContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PcAgent_StopContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PcAgentServer).StopContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.PcAgent/StopContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PcAgentServer).StopContainer(ctx, req.(*StopContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PcAgent_DeleteContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PcAgentServer).DeleteContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.PcAgent/DeleteContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PcAgentServer).DeleteContainer(ctx, req.(*DeleteContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PcAgent_GetStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PcAgentServer).GetStatus(&pcAgentGetStatusServer{stream})
}

type PcAgent_GetStatusServer interface {
	Send(*GetStatsResponse) error
	Recv() (*GetStatusRequest, error)
	grpc.ServerStream
}

type pcAgentGetStatusServer struct {
	grpc.ServerStream
}

func (x *pcAgentGetStatusServer) Send(m *GetStatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pcAgentGetStatusServer) Recv() (*GetStatusRequest, error) {
	m := new(GetStatusRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PcAgent_ServiceDesc is the grpc.ServiceDesc for PcAgent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PcAgent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agent.PcAgent",
	HandlerType: (*PcAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContainerList",
			Handler:    _PcAgent_GetContainerList_Handler,
		},
		{
			MethodName: "CreateContainer",
			Handler:    _PcAgent_CreateContainer_Handler,
		},
		{
			MethodName: "StartContainer",
			Handler:    _PcAgent_StartContainer_Handler,
		},
		{
			MethodName: "RestartContainer",
			Handler:    _PcAgent_RestartContainer_Handler,
		},
		{
			MethodName: "StopContainer",
			Handler:    _PcAgent_StopContainer_Handler,
		},
		{
			MethodName: "DeleteContainer",
			Handler:    _PcAgent_DeleteContainer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStatus",
			Handler:       _PcAgent_GetStatus_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ezdev/agent/pc_agent_server.proto",
}