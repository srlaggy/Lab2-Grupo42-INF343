package sendPlaysDN

import (
	"context"
	sp "lab/datanode/proto/sendPlaysDN"
	dnr "lab/datanode/src/dataNodeRegistro"
	ut "lab/datanode/utils"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	address = "localhost"
	protocolo_grpc = "tcp"
	port_grpc1 = "60101"
	//Revisar
)

// ----- FUNCIÓN: recibir jugadas del Lider ----- // --> NameNode actua como servidor
type server struct {
	sp.UnimplementedJugadasServiceServer
}

// funcion: conecta con el service Jugadas
func (s *server) SendJugadas(ctx context.Context, in *sp.JugadasReq) (*sp.JugadasResp, error) {
	log.Printf("Received %v", in.Registro)
	dnr.RegistrarJugada(in.Registro)
	return &sp.JugadasResp{}, nil
}

// funciones: crea la conexión
func Grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc1)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	sp.RegisterJugadasServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v", port_grpc1)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}