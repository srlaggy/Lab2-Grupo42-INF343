package requestGameLJ

import (
	"context"
	"time"
	"fmt"
	"google.golang.org/grpc"
	ut "lab/jugador/utils"
	lj "lab/jugador/proto/LJ"
)

const (
	address = "localhost"
	protocolo_grpc = ""
	port_grpc = "50000"
	maxJug = 16
)

// variables globales
var nroJugador int64
var nroJugadorBot []int64
var vivosSlice []bool

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: pedir ingresar al juego ----- // --> Jugador es cliente

func RequestGame(entrada string) {
	// Set up a connection to the server.
	conn1, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn1.Close()

	c := lj.NewLiderJugadorServiceClient(conn1)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	nroJugadorBot = make([]int64, maxJug)
	vivosSlice = make([]bool, maxJug)
	for i:=0; i<maxJug; i++{
		vivosSlice[i] = true
	}

	r, err := c.RequestGame(ctx, &lj.GameReq{EntryMsg: entrada})
	ut.FailOnError(err, "Failed to send a entry")
	// almacenamos numero del jugador
	nroJugador = r.GetNroJugador()
	nroJugadorBot[nroJugador-1] = nroJugador
	fmt.Printf("%s", r.GetGameMsg())
}

func RequestGameBot(entrada string) {
	conn1, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn1.Close()

	c := lj.NewLiderJugadorServiceClient(conn1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.RequestGame(ctx, &lj.GameReq{EntryMsg: entrada})
	ut.FailOnError(err, "Failed to send a entry")
	// almacenamos numero del jugador
	nroJugadorBot[r.GetNroJugador() - 1] = r.GetNroJugador()
	// fmt.Printf("%s", r.GetGameMsg())
}

func GetMaxJug() int{
	return maxJug
}

func GetNumeroJugador() int64{
	return nroJugador
}

func GetNumeroJugadorBot() []int64{
	return nroJugadorBot
}

func GetVivosSlice() []bool{
	return vivosSlice
}