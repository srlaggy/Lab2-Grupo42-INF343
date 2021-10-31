package main

// LIDER

import (
	"context"
	"time"
	"fmt"
	"github.com/streadway/amqp"
	"strconv"
	"google.golang.org/grpc"
	pb "lab/lider/grpc"
	ut "lab/lider/utils"
)

const (
	protocolo_rabbit = "amqp"
	address = "localhost"
	port_rabbit = "5672"
	protocolo_grpc = ""
	port_grpc1 = "60000"
	port_grpc2 = "70000"
)

// --------------- FUNCIONES RABBITMQ --------------- //

// ------ ACCIÓN: enviar los jugadores eliminados al pozo ----- //

func muertos(nroJugador int, nroRonda int) string{
	jugador := strconv.Itoa(nroJugador)
	ronda := strconv.Itoa(nroRonda)
	return "Jugador_" + jugador + " Ronda_" + ronda + " "
}

func amqp_func() {
	// Creamos conexion conn
	conn2, err := amqp.Dial(ut.CreateDir(protocolo_rabbit, address, port_rabbit))
	ut.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn2.Close()

	// Creamos canal ch
	ch, err := conn2.Channel()
	ut.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declaramos la cola q
	q, err := ch.QueueDeclare(
		"panaRabbit", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	ut.FailOnError(err, "Failed to declare a queue")

	body := muertos(16,8)

	// routine
	err = ch.Publish(
	"",     // exchange
	q.Name, // routing key
	false,  // mandatory
	false,  // immediate
	amqp.Publishing {
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	ut.FailOnError(err, "Failed to publish a message")
}

// --------------- FUNCIONES GRPC --------------- //

// ----- ACCIÓN: pedir monto acumulado al pozo ----- //

func grpc_func() {
	// Set up a connection to the server.
	conn1, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc1), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn1.Close()

	c := pb.NewPozoServiceClient(conn1)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendMount(ctx, &pb.MountReq{})
	ut.FailOnError(err, "Failed to send a mount")
	fmt.Println("El pozo tiene un monto de: %d", r.GetMonto())
}

// ----- ACCIÓN: enviar jugadas al NameNode ----- //
func enviarJugadasNameNode(jugada string) {
	// // Set up a connection to the server.
	// conn3, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc2), grpc.WithInsecure(), grpc.WithBlock())
	// ut.FailOnError(err, "Failed to create a connection")
	// defer conn3.Close()

	// c := pb.NewPozoServiceClient(conn3)
	// // Contact the server and print out its response.
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// r, err := c.SendMount(ctx, &pb.MountReq{})
	// ut.FailOnError(err, "Failed to send a mount")
	// fmt.Println("El pozo tiene un monto de: %d", r.GetMonto())

}




// --------------- FUNCION MAIN --------------- //

func main(){

	// ----- GRPC ----- //
	grpc_func()

	// ----- RABBITMQ ----- //
	amqp_func()
}