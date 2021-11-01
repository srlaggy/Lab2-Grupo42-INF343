package main

// namenode

import(
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	sp "lab/namenode/proto/sendPlaysNL"
	ut "lab/namenode/utils"
)

const (
	address = "localhost"
	protocolo_grpc = "tcp"
	port_grpc1 = "60001"
)

// --------------- FUNCIONES GRPC --------------- //
type server struct {
	sp.UnimplementedSendPlaysServiceServer
}

func (s *server) SendPlays(ctx context.Context, in *sp.PlaysReq) (*sp.PlaysResp, error) {
	log.Printf("Received %v", in.Play) // 
	return &sp.PlaysResp{Resp: "recibido Hello WORD"}, nil
}

// --------------- FUNCION PRINCIPAL --------------- //
func grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc1)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	sp.RegisterSendPlaysServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v", port_grpc1)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}

func main() {
	grpc_func()
}