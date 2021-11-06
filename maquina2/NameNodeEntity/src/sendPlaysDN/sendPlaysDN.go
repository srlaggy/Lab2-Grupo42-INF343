package sendPlaysDN

import (
	"context"
	sp "lab/namenode/proto/sendPlaysDN"
	ut "lab/namenode/utils"
	"time"
	// "fmt"
	"google.golang.org/grpc"
)

const (
	address1 = "10.6.43.45"
	address2 = "10.6.43.47"
	address3 = "10.6.43.48"
	protocolo_grpc = ""
)

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: enviar jugadas al datanode ----- // --> NameNode actua como cliente
func SendPlaysDataNode(jugada string, ip_datanode string) {
	// fmt.Println(jugada, "su datanode asignado es:", datanode)
	// Creamos conexion
	if(ip_datanode == address1){
		conn3, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address1, "60101"), grpc.WithInsecure(), grpc.WithBlock())
		ut.FailOnError(err, "Failed to create a connection")
			// Creamos conexion con el servicio 
		csp := sp.NewJugadasServiceClient(conn3)
		// Conectamos con el servidor y se imprime la respuesta
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_,err2 := csp.SendJugadas(ctx, &sp.JugadasReq{Registro: jugada})
		ut.FailOnError(err2, "Failed to send a play")
		defer conn3.Close()
	}else if(ip_datanode == address2){
		conn3, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address2, "60103"), grpc.WithInsecure(), grpc.WithBlock())
		ut.FailOnError(err, "Failed to create a connection")
			// Creamos conexion con el servicio 
		csp := sp.NewJugadasServiceClient(conn3)
		// Conectamos con el servidor y se imprime la respuesta
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_,err2 := csp.SendJugadas(ctx, &sp.JugadasReq{Registro: jugada})
		ut.FailOnError(err2, "Failed to send a play")
		defer conn3.Close()
	}else if(ip_datanode == address3){
		conn3, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address3, "60104"), grpc.WithInsecure(), grpc.WithBlock())
		ut.FailOnError(err, "Failed to create a connection")
			// Creamos conexion con el servicio 
		csp := sp.NewJugadasServiceClient(conn3)
		// Conectamos con el servidor y se imprime la respuesta
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_,err2 := csp.SendJugadas(ctx, &sp.JugadasReq{Registro: jugada})
		ut.FailOnError(err2, "Failed to send a play")
		defer conn3.Close()
	}
}