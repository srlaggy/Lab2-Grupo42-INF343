package sendPlaysLJ

import (
	"context"
	"time"
	"fmt"
	"google.golang.org/grpc"
	ut "lab/jugador/utils"
	sp "lab/jugador/proto/sendPlaysLJ"
	e1 "lab/jugador/etapa/e1"
)

const (
	address = "localhost"
	protocolo_grpc = ""
	port_grpc = "50001"
)

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: enviar jugadas al Lider ----- // --> Jugador actua como cliente
func SendPasos() {
	// Creamos conexion
	conn3, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn3.Close()

	// Creamos conexion con el servicio 
	csp := sp.NewSendPasosServiceClient(conn3)
	// Conectamos con el lider y se imprime la respuesta
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	aux := e1.RandomNumber()

	respuesta,err2 := csp.SendPasos(ctx, &sp.NumPasosReq{PlayMsg: aux})
	ut.FailOnError(err2, "Failed to send a play")

	// muerto 0 - vivo 1
	fmt.Printf("tu estas %v, tu numero fue %d", respuesta.GetStateMsg(), aux)
}