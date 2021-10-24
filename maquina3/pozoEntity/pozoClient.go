package main

import (
	"context"
	"log"
	"time"
	"fmt"

	"google.golang.org/grpc"
	pb "lab/pozo/pozo"
)

const (
	address = "localhost:60000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPozoServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendMount(ctx, &pb.MountReq{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println("El pozo tiene un monto de: %d", r.GetMonto())
}
