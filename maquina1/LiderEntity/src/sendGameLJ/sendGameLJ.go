package sendGameLJ

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	rg "lab/lider/proto/requestGameLJ"
	ut "lab/lider/utils"
)

const (
	protocolo_grpc = "tcp"
	port_grpc = "50000"
	mensajeDeEntrada = "Entraste al juego"
)

// --------------- FUNCIONES GRPC --------------- //
type server struct {
	rg.UnimplementedRequestGameServiceServer
}

func (s *server) RequestGame(ctx context.Context, in *rg.GameReq) (*rg.GameResp, error) {
	log.Printf("Entry Received")
	return &rg.GameResp{GameMsg: mensajeDeEntrada}, nil
}

// --------------- FUNCION PRINCIPAL --------------- //
func Grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	rg.RegisterRequestGameServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v", port_grpc)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}