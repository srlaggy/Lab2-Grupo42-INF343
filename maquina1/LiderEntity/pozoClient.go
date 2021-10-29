package main

// LIDER

import (
	"context"
	"time"
	"fmt"

	"google.golang.org/grpc"
	pb "lab/lider/grpc"
)

const (
	protocolo = ""
	address = "localhost"
	port = "60000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(createDir(protocolo, address, port), grpc.WithInsecure(), grpc.WithBlock())
	failOnError(err, "Failed to create a connection")
	defer conn.Close()

	c := pb.NewPozoServiceClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendMount(ctx, &pb.MountReq{})
	failOnError(err, "Failed to send a mount")
	fmt.Println("El pozo tiene un monto de: %d", r.GetMonto())
}