// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// CServiceClient is the client API for CService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CServiceClient interface {
	EchoC(ctx context.Context, in *CMsg, opts ...grpc.CallOption) (*CMsg, error)
}

type cServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCServiceClient(cc grpc.ClientConnInterface) CServiceClient {
	return &cServiceClient{cc}
}

func (c *cServiceClient) EchoC(ctx context.Context, in *CMsg, opts ...grpc.CallOption) (*CMsg, error) {
	out := new(CMsg)
	err := c.cc.Invoke(ctx, "/pb.CService/EchoC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CServiceServer is the server API for CService service.
// All implementations must embed UnimplementedCServiceServer
// for forward compatibility
type CServiceServer interface {
	EchoC(context.Context, *CMsg) (*CMsg, error)
	mustEmbedUnimplementedCServiceServer()
}

// UnimplementedCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCServiceServer struct {
}

func (UnimplementedCServiceServer) EchoC(context.Context, *CMsg) (*CMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoC not implemented")
}
func (UnimplementedCServiceServer) mustEmbedUnimplementedCServiceServer() {}

// UnsafeCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CServiceServer will
// result in compilation errors.
type UnsafeCServiceServer interface {
	mustEmbedUnimplementedCServiceServer()
}

func RegisterCServiceServer(s grpc.ServiceRegistrar, srv CServiceServer) {
	s.RegisterService(&_CService_serviceDesc, srv)
}

func _CService_EchoC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CServiceServer).EchoC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CService/EchoC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CServiceServer).EchoC(ctx, req.(*CMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _CService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CService",
	HandlerType: (*CServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EchoC",
			Handler:    _CService_EchoC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "c.proto",
}
