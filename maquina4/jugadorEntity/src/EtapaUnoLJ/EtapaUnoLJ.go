package sendPlaysLJ

import (
	"context"
	"time"
	"fmt"
	"google.golang.org/grpc"
	ut "lab/jugador/utils"
	lj "lab/jugador/proto/LJ"
	e1 "lab/jugador/etapa/e1"
	rg "lab/jugador/src/requestGameLJ"
)

const (
	address = "localhost"
	protocolo_grpc = ""
	port_grpc = "50001"
)

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: enviar jugadas al Lider ----- // --> Jugador actua como cliente
func EtapaUno() {
	// Creamos conexion
	conn3, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection\n")
	defer conn3.Close()

	// Creamos conexion con el servicio 
	csp := lj.NewLiderJugadorServiceClient(conn3)
	// Conectamos con el lider y se imprime la respuesta
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	aux := e1.RandomNumber()

	respuesta,err2 := csp.EtapaUno(ctx, &lj.NumPasosReq{PlayMsg: aux, NroJugador: rg.GetNumeroJugador()})
	ut.FailOnError(err2, "Failed to send a play\n")

	// muerto 0 - vivo 1
	e1.LiveOrDead(respuesta.GetStateMsg())
	fmt.Printf("Tu numero fue %d\n", aux)
}