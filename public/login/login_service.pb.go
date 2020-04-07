// Code generated by protoc-gen-go. DO NOT EDIT.
// source: login_service.proto

package Login

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("login_service.proto", fileDescriptor_565c413639ec516a) }

var fileDescriptor_565c413639ec516a = []byte{
	// 104 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0xc9, 0x4f, 0xcf,
	0xcc, 0x8b, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0xf5, 0x01, 0x09, 0x4a, 0xf1, 0x24, 0xe7, 0x64, 0xa6, 0xe6, 0x95, 0x40, 0x04, 0x8d, 0x1c,
	0xb8, 0x38, 0x1c, 0x0b, 0x0a, 0xc0, 0x32, 0x42, 0x26, 0x5c, 0x10, 0x25, 0x42, 0xc2, 0x7a, 0x60,
	0x1a, 0x42, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x48, 0x89, 0xa0, 0x0a, 0x16, 0x17, 0xe4,
	0xe7, 0x15, 0xa7, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0x0d, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x8e, 0x01, 0x6b, 0x6d, 0x74, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AppLoginClient is the client API for AppLogin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AppLoginClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type appLoginClient struct {
	cc *grpc.ClientConn
}

func NewAppLoginClient(cc *grpc.ClientConn) AppLoginClient {
	return &appLoginClient{cc}
}

func (c *appLoginClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/Login.AppLogin/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppLoginServer is the server API for AppLogin service.
type AppLoginServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

// UnimplementedAppLoginServer can be embedded to have forward compatible implementations.
type UnimplementedAppLoginServer struct {
}

func (*UnimplementedAppLoginServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterAppLoginServer(s *grpc.Server, srv AppLoginServer) {
	s.RegisterService(&_AppLogin_serviceDesc, srv)
}

func _AppLogin_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppLoginServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Login.AppLogin/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppLoginServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppLogin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Login.AppLogin",
	HandlerType: (*AppLoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AppLogin_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "login_service.proto",
}
