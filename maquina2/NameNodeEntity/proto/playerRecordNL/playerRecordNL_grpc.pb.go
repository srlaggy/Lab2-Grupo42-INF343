// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package namenode

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

// PlayerRecordServiceClient is the client API for PlayerRecordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlayerRecordServiceClient interface {
	PlayerRecord(ctx context.Context, in *PlayerReq, opts ...grpc.CallOption) (*PlayerResp, error)
}

type playerRecordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayerRecordServiceClient(cc grpc.ClientConnInterface) PlayerRecordServiceClient {
	return &playerRecordServiceClient{cc}
}

func (c *playerRecordServiceClient) PlayerRecord(ctx context.Context, in *PlayerReq, opts ...grpc.CallOption) (*PlayerResp, error) {
	out := new(PlayerResp)
	err := c.cc.Invoke(ctx, "/playerRecordNL.PlayerRecordService/PlayerRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerRecordServiceServer is the server API for PlayerRecordService service.
// All implementations must embed UnimplementedPlayerRecordServiceServer
// for forward compatibility
type PlayerRecordServiceServer interface {
	PlayerRecord(context.Context, *PlayerReq) (*PlayerResp, error)
	mustEmbedUnimplementedPlayerRecordServiceServer()
}

// UnimplementedPlayerRecordServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPlayerRecordServiceServer struct {
}

func (UnimplementedPlayerRecordServiceServer) PlayerRecord(context.Context, *PlayerReq) (*PlayerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerRecord not implemented")
}
func (UnimplementedPlayerRecordServiceServer) mustEmbedUnimplementedPlayerRecordServiceServer() {}

// UnsafePlayerRecordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlayerRecordServiceServer will
// result in compilation errors.
type UnsafePlayerRecordServiceServer interface {
	mustEmbedUnimplementedPlayerRecordServiceServer()
}

func RegisterPlayerRecordServiceServer(s grpc.ServiceRegistrar, srv PlayerRecordServiceServer) {
	s.RegisterService(&PlayerRecordService_ServiceDesc, srv)
}

func _PlayerRecordService_PlayerRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerRecordServiceServer).PlayerRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playerRecordNL.PlayerRecordService/PlayerRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerRecordServiceServer).PlayerRecord(ctx, req.(*PlayerReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PlayerRecordService_ServiceDesc is the grpc.ServiceDesc for PlayerRecordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlayerRecordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "playerRecordNL.PlayerRecordService",
	HandlerType: (*PlayerRecordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlayerRecord",
			Handler:    _PlayerRecordService_PlayerRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "playerRecordNL.proto",
}
