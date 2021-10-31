package sendMountPL

import (
	"context"
	"log"
	"net"
	"os"
	"bufio"
	"strings"
	"strconv"

	"google.golang.org/grpc"
	sm "lab/pozo/proto/sendMountPL"
	ut "lab/pozo/utils"
)

const (
	protocolo_grpc = "tcp"
	port_grpc = "60000"
	fileName = "utils/pozo.txt"
)

// --------------- FUNCIONES GRPC --------------- //
type server struct {
	sm.UnimplementedPozoServiceServer
}

func (s *server) SendMount(ctx context.Context, in *sm.MountReq) (*sm.MountResp, error) {
	log.Printf("Received")
	return &sm.MountResp{Monto: GetMountPozo()}, nil
}

// Funcion que retorna el monto acumulado del pozo
func GetMountPozo() float64{
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
	mountInt,err := strconv.ParseFloat(mount, 64)
	ut.FailOnError(err, "Failed to transform to int")
	return mountInt
}

// --------------- FUNCION PRINCIPAL --------------- //
func Grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	sm.RegisterPozoServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v", port_grpc)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}