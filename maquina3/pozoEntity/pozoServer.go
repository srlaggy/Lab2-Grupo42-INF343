package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "lab/pozo/pozo"
)

const (
	port = ":60000"
)

type server struct {
	pb.UnimplementedPozoServiceServer
}

func (s *server) SendMount(ctx context.Context, in *pb.MountReq) (*pb.MountResp, error) {
	log.Printf("Received")
	return &pb.MountResp{Monto: 1234567890}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPozoServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}