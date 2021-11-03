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
	max = 10
)

func RandomNumber() int64{
	rand.Seed(time.Now().UnixNano())
	aux := rand.Intn(max - min) + min
	return int64(aux)
}

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: enviar jugadas al Lider ----- // --> Jugador actua como cliente
func EtapaUno(numero int64, contador int64, ronda int64) int64{
	// Creamos conexion
	conn3, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection\n")
	defer conn3.Close()

	// Creamos conexion con el servicio 
	csp := lj.NewLiderJugadorServiceClient(conn3)
	// Conectamos con el lider y se imprime la respuesta
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	respuesta,err2 := csp.EtapaUno(ctx, &lj.NumPasosReq{PlayMsg: numero, NroJugador: rg.GetNumeroJugador(), Total: contador, Ronda: ronda})
	ut.FailOnError(err2, "Failed to send a play\n")

	return respuesta.GetStateMsg()
}

// Juego 1 por rondas (4 rondas maximo)
func StartGameUno(){
	var sumador int64 = 0
	fmt.Println("\nInicio Juego 1")
	for i:=0; i<4; i++{
		fmt.Println("\nInicio Ronda", i+1)
		aux := RandomNumber()
		resp := EtapaUno(aux, sumador, int64(i))
		if resp==0{
			fmt.Println("Tu estas muerto, BANG!!")
			fmt.Println("Tu numero fue", aux)
			break
		}
		sumador += aux
		if i==3{
			if sumador<21 || resp==2{
				fmt.Println("No lograste llegar al final, tu estas muerto, BANG!!")
				fmt.Println("Tu numero fue", aux)
				fmt.Println("Tu total fue", sumador)
				break
			} else {
				fmt.Println("Superaste el juego 1, Felicidades!!")
				fmt.Println("Tu numero fue", aux)
				fmt.Println("Tu total fue", sumador)
				break
			}
		}
		fmt.Println("Te salvaste esta vez")
		fmt.Println("Tu numero fue", aux)
		fmt.Println("Tu total acumulado es de", sumador)
		time.Sleep(3*time.Second)
	}
}