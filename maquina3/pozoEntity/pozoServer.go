package main

// POZO

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "lab/pozo/pozo"
)

const (
	protocolo = "tcp"
	address = ""
	port = "60000"
)

type server struct {
	pb.UnimplementedPozoServiceServer
}

func (s *server) SendMount(ctx context.Context, in *pb.MountReq) (*pb.MountResp, error) {
	log.Printf("Received")
	return &pb.MountResp{Monto: 1234567890}, nil
}

func main() {
	lis, err := net.Listen(protocolo, ":"+port)
	failOnError(err, "Failed to listen")

	s := grpc.NewServer()
	pb.RegisterPozoServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	failOnError(s.Serve(lis), "Failed to serve")
}