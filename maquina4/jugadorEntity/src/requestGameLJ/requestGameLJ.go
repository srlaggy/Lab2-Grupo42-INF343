package requestGameLJ

import (
	"context"
	"time"
	"fmt"
	"google.golang.org/grpc"
	ut "lab/jugador/utils"
	rg "lab/jugador/proto/requestGameLJ"
)

const (
	address = "localhost"
	protocolo_grpc = ""
	port_grpc = "50000"
)

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: pedir ingresar al juego ----- // --> Jugador es cliente

func RequestGame(entrada string) {
	// Set up a connection to the server.
	conn1, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn1.Close()

	c := rg.NewRequestGameServiceClient(conn1)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.RequestGame(ctx, &rg.GameReq{EntryMsg: entrada})
	ut.FailOnError(err, "Failed to send a entry")
	fmt.Printf("%s", r.GetGameMsg())
}