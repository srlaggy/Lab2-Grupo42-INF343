package sendPlaysDN

import (
	"context"
	sp "lab/namenode/proto/sendPlaysDN"
	ut "lab/namenode/utils"
	"time"

	"google.golang.org/grpc"
)

const (
	protocolo_grpc = ""
	port_grpc = "60101"
)

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: enviar jugadas al datanode ----- // --> NameNode actua como cliente
func SendPlaysDataNode(jugada string, address string) {
	// Creamos conexion
	conn3, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn3.Close()

	// Creamos conexion con el servicio 
	csp := sp.NewJugadasServiceClient(conn3)
	// Conectamos con el servidor y se imprime la respuesta
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_,err2 := csp.SendJugadas(ctx, &sp.JugadasReq{Registro: jugada})
	ut.FailOnError(err2, "Failed to send a play")
}