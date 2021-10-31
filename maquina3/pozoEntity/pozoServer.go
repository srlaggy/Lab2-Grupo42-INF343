package main

// POZO

import (
	"context"
	"log"
	"net"
	"os"
	"bufio"
	"strings"
	"strconv"

	"google.golang.org/grpc"
	"github.com/streadway/amqp"
	pb "lab/pozo/pozo"
	ut "lab/pozo/utils"
)

const (
	protocolo_grpc = "tcp"
	protocolo_amqp = "amqp"
	address = "localhost"
	port_grpc = "60000"
	port_amqp = "5672"
	montoMuerto = 1000000
	fileName = "utils/pozo.txt"
)

// --------------- FUNCIONES GRPC --------------- //
type server struct {
	pb.UnimplementedPozoServiceServer
}

func (s *server) SendMount(ctx context.Context, in *pb.MountReq) (*pb.MountResp, error) {
	log.Printf("Received")
	return &pb.MountResp{Monto: 1234567890}, nil
}

func grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	pb.RegisterPozoServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v", port_grpc)
	ut.FailOnError(s.Serve(lis), "Failed to serve")

}

// --------------- FUNCIONES RABBITMQ --------------- //
// Funcion que agrega un muerto al pozo
func writePozo(dead string){

	// Recuperamos monto total del pozo
	monto := getMountPozo()
	newMonto := monto + montoMuerto

	// Abrimos archivo en modo append
    f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	ut.FailOnError(err, "Failed to open file")
	defer f.Close()
	
	// string formateado
	newDead := dead + strconv.Itoa(newMonto) + "\n"

	// Escribimos en archivo
	_, err1 := f.WriteString(newDead)
	ut.FailOnError(err1, "Failed to write file")
}

// Funcion que retorna el monto acumulado del pozo
func getMountPozo() int{
	f, err := os.Open(fileName)
	ut.FailOnError(err, "Failed to open file")
    defer f.Close()

    scanner := bufio.NewScanner(f)
	var auxLine string
    for scanner.Scan(){
		auxLine = scanner.Text()
    }
	ut.FailOnError(scanner.Err(), "Failed to read file")

	mount := strings.Split(auxLine, " ")[2]
	mountInt,err := strconv.Atoi(mount)
	ut.FailOnError(err, "Failed to transform to int")
	return mountInt
}

// Funcion que reinicia el pozo con monto 0
func resetPozo(){
	f, err := os.Create(fileName)
	ut.FailOnError(err, "Failed to create file")
    defer f.Close()
	
	_, err2 := f.WriteString("Jugador_0 Ronda_0 0\n")
	ut.FailOnError(err2, "Failed to write file")
}

func amqp_func(){
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


// --------------- FUNCION MAIN --------------- //

func main() {
    // ----- GRPC ----- //
    go grpc_func()

    // ----- RABBITMQ ----- //
    amqp_func()
}