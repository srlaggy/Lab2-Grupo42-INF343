package main

import (
	"github.com/streadway/amqp"
	"strconv"
)

// LIDER

const (
	protocolo = "amqp"
	address = "localhost"
	port = "5672"
)

func muertos(nroJugador int, nroRonda int) string{
	jugador := strconv.Itoa(nroJugador)
	ronda := strconv.Itoa(nroRonda)
	return "Jugador_" + jugador + " Ronda_" + ronda + " "
}

func main(){
	// Creamos conexion conn
	conn, err := amqp.Dial(createDir(protocolo, address, port))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Creamos canal ch
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
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
	failOnError(err, "Failed to declare a queue")

	body := muertos(10,4)

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
	failOnError(err, "Failed to publish a message")
}