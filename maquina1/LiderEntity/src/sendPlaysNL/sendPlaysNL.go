package sendPlaysNL

import (
	"context"
	"time"
	"google.golang.org/grpc"
	ut "lab/lider/utils"
	sp "lab/lider/proto/sendPlaysNL"
)

const (
	address = "localhost"
	protocolo_grpc = ""
	port_grpc = "60001"
)

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: enviar jugadas al NameNode ----- // --> Lider actua como cliente
func SendPlaysLider(jugada string) {
	// Creamos conexion
	conn3, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn3.Close()

	// Creamos conexion con el servicio 
	csp := sp.NewSendPlaysServiceClient(conn3)
	// Conectamos con el servidor y se imprime la respuesta
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_,err2 := csp.SendPlays(ctx, &sp.PlaysReq{Play: jugada})
	ut.FailOnError(err2, "Failed to send a play")
}