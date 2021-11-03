// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package nameNode

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

// RecordServiceClient is the client API for RecordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RecordServiceClient interface {
	SendRecord(ctx context.Context, in *RecordReq, opts ...grpc.CallOption) (*RecordResp, error)
}

type recordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRecordServiceClient(cc grpc.ClientConnInterface) RecordServiceClient {
	return &recordServiceClient{cc}
}

func (c *recordServiceClient) SendRecord(ctx context.Context, in *RecordReq, opts ...grpc.CallOption) (*RecordResp, error) {
	out := new(RecordResp)
	err := c.cc.Invoke(ctx, "/playerRecordDN.RecordService/SendRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecordServiceServer is the server API for RecordService service.
// All implementations must embed UnimplementedRecordServiceServer
// for forward compatibility
type RecordServiceServer interface {
	SendRecord(context.Context, *RecordReq) (*RecordResp, error)
	mustEmbedUnimplementedRecordServiceServer()
}

// UnimplementedRecordServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRecordServiceServer struct {
}

func (UnimplementedRecordServiceServer) SendRecord(context.Context, *RecordReq) (*RecordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRecord not implemented")
}
func (UnimplementedRecordServiceServer) mustEmbedUnimplementedRecordServiceServer() {}

// UnsafeRecordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RecordServiceServer will
// result in compilation errors.
type UnsafeRecordServiceServer interface {
	mustEmbedUnimplementedRecordServiceServer()
}

func RegisterRecordServiceServer(s grpc.ServiceRegistrar, srv RecordServiceServer) {
	s.RegisterService(&RecordService_ServiceDesc, srv)
}

func _RecordService_SendRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceServer).SendRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playerRecordDN.RecordService/SendRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceServer).SendRecord(ctx, req.(*RecordReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RecordService_ServiceDesc is the grpc.ServiceDesc for RecordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RecordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "playerRecordDN.RecordService",
	HandlerType: (*RecordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRecord",
			Handler:    _RecordService_SendRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "playerRecordDN.proto",
}