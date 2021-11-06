package sendDeadPL

import (
	"github.com/streadway/amqp"
	"strconv"
	ut "lab/lider/utils"
)

const (
	protocolo_rabbit = "amqp"
	address_rabbit = "localhost"
	port_rabbit = "5672"
)

// --------------- FUNCIONES RABBITMQ --------------- //

// ------ FUNCIÓN: enviar los jugadores eliminados al pozo ----- // --> Lider actua como productor

func Muertos(nroJugador int, nroRonda int) string{
	jugador := strconv.Itoa(nroJugador)
	ronda := strconv.Itoa(nroRonda)
	return "Jugador_" + jugador + " Ronda_" + ronda + " "
}

func SendDead_amqp(body string) {
	// Creamos conexion conn
	conn2, err := amqp.Dial(ut.CreateDir(protocolo_rabbit, address_rabbit, port_rabbit))
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