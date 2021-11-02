// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package lider

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

// SendPasosServiceClient is the client API for SendPasosService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendPasosServiceClient interface {
	SendPasos(ctx context.Context, in *NumPasosReq, opts ...grpc.CallOption) (*NumPasosResp, error)
}

type sendPasosServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSendPasosServiceClient(cc grpc.ClientConnInterface) SendPasosServiceClient {
	return &sendPasosServiceClient{cc}
}

func (c *sendPasosServiceClient) SendPasos(ctx context.Context, in *NumPasosReq, opts ...grpc.CallOption) (*NumPasosResp, error) {
	out := new(NumPasosResp)
	err := c.cc.Invoke(ctx, "/sendPlaysLJ.SendPasosService/SendPasos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendPasosServiceServer is the server API for SendPasosService service.
// All implementations must embed UnimplementedSendPasosServiceServer
// for forward compatibility
type SendPasosServiceServer interface {
	SendPasos(context.Context, *NumPasosReq) (*NumPasosResp, error)
	mustEmbedUnimplementedSendPasosServiceServer()
}

// UnimplementedSendPasosServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSendPasosServiceServer struct {
}

func (UnimplementedSendPasosServiceServer) SendPasos(context.Context, *NumPasosReq) (*NumPasosResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPasos not implemented")
}
func (UnimplementedSendPasosServiceServer) mustEmbedUnimplementedSendPasosServiceServer() {}

// UnsafeSendPasosServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SendPasosServiceServer will
// result in compilation errors.
type UnsafeSendPasosServiceServer interface {
	mustEmbedUnimplementedSendPasosServiceServer()
}

func RegisterSendPasosServiceServer(s grpc.ServiceRegistrar, srv SendPasosServiceServer) {
	s.RegisterService(&SendPasosService_ServiceDesc, srv)
}

func _SendPasosService_SendPasos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumPasosReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendPasosServiceServer).SendPasos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sendPlaysLJ.SendPasosService/SendPasos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendPasosServiceServer).SendPasos(ctx, req.(*NumPasosReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SendPasosService_ServiceDesc is the grpc.ServiceDesc for SendPasosService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendPasosService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sendPlaysLJ.SendPasosService",
	HandlerType: (*SendPasosServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPasos",
			Handler:    _SendPasosService_SendPasos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sendPlaysLJ.proto",
}
