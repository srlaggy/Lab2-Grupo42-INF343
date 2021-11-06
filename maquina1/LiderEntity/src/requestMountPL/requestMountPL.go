package requestMountPL

import (
	"context"
	"time"
	// "fmt"
	"google.golang.org/grpc"
	ut "lab/lider/utils"
	sm "lab/lider/proto/sendMountPL"
)

const (
	address = "10.6.43.47"
	protocolo_grpc = ""
	port_grpc = "60000"
)

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: pedir monto acumulado al pozo ----- // --> Lider actua como cliente

func RequestMount() int64{
	// Set up a connection to the server.
	conn1, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn1.Close()

	c := sm.NewPozoServiceClient(conn1)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendMount(ctx, &sm.MountReq{})
	ut.FailOnError(err, "Failed to send a mount")
	// fmt.Printf("El pozo tiene un monto de: %.f", r.GetMonto())
	return int64(r.GetMonto())
}