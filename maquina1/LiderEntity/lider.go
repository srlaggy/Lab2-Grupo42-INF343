package main

// LIDER

import (
	"context"
	"time"
	"fmt"
	"github.com/streadway/amqp"
	"strconv"
	"google.golang.org/grpc"
	rm "lab/lider/grpc"
	ut "lab/lider/utils"
	sp "lab/lider/grpc/sendPlaysNL"
)

const (
	protocolo_rabbit = "amqp"
	address = "localhost"
	port_rabbit = "5672"
	protocolo_grpc = ""
	port_grpc1 = "60000"
	port_grpc2 = "60001"
)

// --------------- FUNCIONES RABBITMQ --------------- //

// ------ FUNCIÓN: enviar los jugadores eliminados al pozo ----- // --> Lider actua como productor

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

// ----- FUNCIÓN: pedir monto acumulado al pozo ----- // --> Lider actua como cliente

func requestMount() {
	// Set up a connection to the server.
	conn1, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc1), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn1.Close()

	c := rm.NewPozoServiceClient(conn1)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendMount(ctx, &rm.MountReq{})
	ut.FailOnError(err, "Failed to send a mount")
	fmt.Printf("El pozo tiene un monto de: %.f", r.GetMonto())
}

// ----- FUNCIÓN: enviar jugadas al NameNode ----- // --> Lider actua como cliente
func sendPlaysLider(jugada string) {
	// Creamos conexion
	conn3, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc2), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn3.Close()

	// Creamos conexion con el servicio 
	csp := sp.NewSendPlaysServiceClient(conn3)
	// Conectamos con el servidor y se imprime la respuesta
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := csp.SendPlays(ctx, &sp.PlaysReq{Play: jugada}) // 
	ut.FailOnError(err, "Failed to send a play")
	fmt.Println("Registro exitoso", r)
}




// --------------- FUNCION MAIN --------------- //

func main(){

	// ----- FUNCIÓN: pedir monto acumulado al pozo ----- //
	// go requestMount()

	// ----- FUNCIÓN: enviar jugadas al NameNode ----- //
	go sendPlaysLider("Jugador_1 Ronda_1 jugada_1") 	// 
	// time.Sleep(5 * time.Second)
	// ----- FUNCIÓN: enviar los jugadores eliminados al pozo ----- //
	amqp_func()
}