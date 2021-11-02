package receivePlaysLJ

import(
	"context"
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
	lj "lab/lider/proto/LJ"
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
	lj.UnimplementedLiderJugadorServiceServer
}

// funcion: primer juego
func (s *server) EtapaUno(ctx context.Context, in *lj.NumPasosReq) (*lj.NumPasosResp, error) {
	var msg int64 = 1
	if in.PlayMsg >= nroLider{
		fmt.Printf("El jugador %d ha muerto\n", in.NroJugador)
		msg = 0
	}
	return &lj.NumPasosResp{StateMsg: msg}, nil
}

// funciones: crea la conexión
func Grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc1)
	ut.FailOnError(err, "Failed to listen")

	fmt.Printf("el numero del lider es %d\n", nroLider)

	s := grpc.NewServer()
	lj.RegisterLiderJugadorServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v", port_grpc1)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}