// Funcion que reinicia el pozo con monto 0
// func resetPozo(){
// 	f, err := os.Create(fileName)
// 	ut.FailOnError(err, "Failed to create file")
//     defer f.Close()
// 	_, err2 := f.WriteString("Jugador_0 Ronda_0 0\n")
// 	ut.FailOnError(err2, "Failed to write file")
// }

package receiveDeadPL

import (
	"log"
	"os"
	"fmt"

	"github.com/streadway/amqp"
	ut "lab/pozo/utils"
	sm "lab/pozo/src/sendMountPL"
)

const (
	protocolo_amqp = "amqp"
	address = "localhost"
	port_amqp = "5672"
	montoMuerto = 100000000
	fileName = "utils/pozo.txt"
)

// --------------- FUNCIONES RABBITMQ --------------- //
// Funcion que agrega un muerto al pozo
func writePozo(dead string){

	// Recuperamos monto total del pozo
	monto := sm.GetMountPozo()
	newMonto := monto + montoMuerto

	// Abrimos archivo en modo append
    f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	ut.FailOnError(err, "Failed to open file")
	defer f.Close()
	
	// string formateado
	newDead := dead + fmt.Sprintf("%.f", newMonto) + "\n"

	// Escribimos en archivo
	_, err1 := f.WriteString(newDead)
	ut.FailOnError(err1, "Failed to write file")
}

// --------------- FUNCION PRINCIPAL --------------- //
func Rabbit_func(){
	// ----- RABBITMQ ----- //
	// Creacion de conexion
	conn, err := amqp.Dial(ut.CreateDir(protocolo_amqp, address, port_amqp))
	ut.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Creacion de canal
	ch, err := conn.Channel()
	ut.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Declaracion de cola
	q, err := ch.QueueDeclare(
		"panaRabbit", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	ut.FailOnError(err, "Failed to declare a queue")

	// Consumiendo canal
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	ut.FailOnError(err, "Failed to register a consumer")
	forever := make(chan bool)

	// routine
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			writePozo(string(d.Body))
		}
	}()
	log.Printf("Servidor rabbitmq escuchando en el puerto %v", port_amqp)
	<-forever
}