package receivePlaysLJ

import(
	"context"
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
	sp "lab/lider/proto/sendPlaysLJ"
	e1 "lab/lider/etapa/e1"
	ut "lab/lider/utils"
)

const (
	address = "localhost"
	protocolo_grpc = "tcp"
	port_grpc1 = "50001"
)

// var global
var nroLider int64 = e1.RandomNumber()

// ----- FUNCIÓN: recibir jugadas del jugador ----- // --> Lider actua como servidor
type server struct {
	sp.UnimplementedSendPasosServiceServer
}

// funcion: conecta con el service SendPasos
func (s *server) SendPasos(ctx context.Context, in *sp.NumPasosReq) (*sp.NumPasosResp, error) {
	var msg int64 = 1
	if in.PlayMsg >= nroLider{
		fmt.Printf("El jugador x ha muerto\n")
		msg = 0
	}
	return &sp.NumPasosResp{StateMsg: msg}, nil
}

// funciones: crea la conexión
func Grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc1)
	ut.FailOnError(err, "Failed to listen")

	fmt.Printf("el numero del lider es %d\n", nroLider)

	s := grpc.NewServer()
	sp.RegisterSendPasosServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v", port_grpc1)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}