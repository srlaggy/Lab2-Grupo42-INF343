package main

import (
	"log"
	"github.com/streadway/amqp"
	"os"
	"bufio"
	"strings"
	"strconv"
)

// POZO

const (
	protocolo = "amqp"
	address = "localhost"
	port = "5672"
	montoMuerto = 1000000
)

// Funcion que agrega un muerto al pozo
func writePozo(dead string){
	fileName := "pozo.txt"

	// Recuperamos monto total del pozo
	monto := getMountPozo()
	newMonto := monto + montoMuerto

	// Abrimos archivo en modo append
    f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	failOnError(err, "Failed to open file")
	defer f.Close()
	
	// string formateado
	newDead := dead + strconv.Itoa(newMonto) + "\n"

	// Escribimos en archivo
	_, err1 := f.WriteString(newDead)
	failOnError(err1, "Failed to write file")
}

// Funcion que retorna el monto acumulado del pozo
func getMountPozo() int{
	f, err := os.Open("pozo.txt")
	failOnError(err, "Failed to open file")
    defer f.Close()

    scanner := bufio.NewScanner(f)
	var auxLine string
    for scanner.Scan(){
		auxLine = scanner.Text()
    }
	failOnError(scanner.Err(), "Failed to read file")

	mount := strings.Split(auxLine, " ")[2]
	mountInt,err := strconv.Atoi(mount)
	failOnError(err, "Failed to transform to int")
	return mountInt
}

// Funcion que reinicia el pozo con monto 0
func resetPozo(){
	f, err := os.Create("pozo.txt")
	failOnError(err, "Failed to create file")
    defer f.Close()
	
	_, err2 := f.WriteString("Jugador_0 Ronda_0 0\n")
	failOnError(err2, "Failed to write file")
}

func main(){
	// Creacion de conexion
	conn, err := amqp.Dial(createDir(protocolo, address, port))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Creacion de canal
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
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
	failOnError(err, "Failed to declare a queue")

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
	failOnError(err, "Failed to register a consumer")
	forever := make(chan bool)

	// routine
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			writePozo(string(d.Body))
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}