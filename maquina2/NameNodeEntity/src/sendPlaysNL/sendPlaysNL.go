package sendPlaysNL

import (
	"context"
	sp "lab/namenode/proto/sendPlaysNL"
	am "lab/namenode/src/administracionJugadas"
	ut "lab/namenode/utils"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	address = "localhost"
	protocolo_grpc = "tcp"
	port_grpc1 = "60001"
)

// ----- FUNCIÓN: recibir jugadas del Lider ----- // --> NameNode actua como servidor
type server struct {
	sp.UnimplementedSendPlaysServiceServer
}

// funcion: conecta con el service SendPlays
func (s *server) SendPlays(ctx context.Context, in *sp.PlaysReq) (*sp.PlaysResp, error) {
	log.Printf("Received %v", in.Play) // 
	am.EntregarJugada(in.Play)
	return &sp.PlaysResp{}, nil
}

// funciones: crea la conexión
func Grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc1)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	sp.RegisterSendPlaysServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v", port_grpc1)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}