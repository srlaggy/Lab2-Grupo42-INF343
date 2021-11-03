package sendGameLJ

import (
	"context"
	"log"
	"net"
	"fmt"

	"google.golang.org/grpc"
	lj "lab/lider/proto/LJ"
	sg "lab/lider/src/startGameL"
	ut "lab/lider/utils"
)

const (
	protocolo_grpc = "tcp"
	port_grpc = "50000"
)

// VARIABLES GLOBALES
var mensajeDeEntrada string

// --------------- FUNCIONES GRPC --------------- //
type server struct {
	lj.UnimplementedLiderJugadorServiceServer
}

func (s *server) RequestGame(ctx context.Context, in *lj.GameReq) (*lj.GameResp, error) {
	value := sg.AddPlayerGame()
	mensajeDeEntrada = "\nEl juego ya comenzó. No puedes ingresar.\n"
	if value!=0{
		mensajeDeEntrada = fmt.Sprintf("\nEstas dentro del juego. Eres el jugador %d\n", value)
		log.Printf("Entry Received")
	}
	return &lj.GameResp{GameMsg: mensajeDeEntrada, NroJugador: value}, nil
}

// --------------- FUNCION PRINCIPAL --------------- //
func Grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	lj.RegisterLiderJugadorServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v\n", port_grpc)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}