package sendPlaysLJ

import (
	"context"
	"time"
	"fmt"
	"math/rand"
	"google.golang.org/grpc"
	ut "lab/jugador/utils"
	lj "lab/jugador/proto/LJ"
	rg "lab/jugador/src/requestGameLJ"
)

const (
	address = "localhost"
	protocolo_grpc = ""
	port_grpc = "50001"
	min = 1
	max = 4
)

var vivo_muerto bool = true

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: Etapa 2 ----- // --> Jugador actua como cliente

func Etapa2Conn() {
	// Creamos conexion
	conn3, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn3.Close()

	// Creamos conexion con el servicio 
	csp := lj.NewLiderJugadorServiceClient(conn3)
	// Conectamos con el lider y se imprime la respuesta
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r,err2 := csp.Etapa2Conn(ctx, &lj.E2ConnReq{NroJugador: rg.GetNumeroJugador()})
	ut.FailOnError(err2, "Failed to send a play")
	if(r.GetNroGroup() == 1){
		fmt.Println("Eres parte del Grupo 1")
	} else if(r.GetNroGroup() == 2){
		fmt.Println("Eres parte del Grupo 2")
	} else{
		vivo_muerto = false
		fmt.Println("Mala suerte, moriste")
	}
}

// determina numero random del jugador
func RandomNumber() int64{
	rand.Seed(time.Now().UnixNano())
	aux := rand.Intn(max - min) + min
	return int64(aux)
}

func Etapa2() {
	// Creamos conexion
	conn3, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn3.Close()

	// Creamos conexion con el servicio 
	csp := lj.NewLiderJugadorServiceClient(conn3)
	// Conectamos con el lider y se imprime la respuesta
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r,err2 := csp.Etapa2(ctx, &lj.Etapa2Req{NroJugador: rg.GetNumeroJugador(), Numero: RandomNumber()})
	ut.FailOnError(err2, "Failed to send a play")
	fmt.Println(r.GetStateMsg())
}

func GetVivo_muerto() bool{
	return vivo_muerto
}