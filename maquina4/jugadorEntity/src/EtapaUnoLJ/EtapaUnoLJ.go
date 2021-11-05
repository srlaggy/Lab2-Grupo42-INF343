package sendPlaysLJ

import (
	"context"
	"time"
	"fmt"
	"sync"
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

// determina numero random del jugador
func RandomNumber(num int64) int64{
	rand.Seed(time.Now().UnixNano() + num*5)
	aux := rand.Intn(max - min) + min
	return int64(aux)
}

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: enviar jugadas al Lider ----- // --> Jugador actua como cliente
func EtapaUno(numero int64, contador int64, ronda int64, nroJugadorAux int64) int64{
	// Creamos conexion
	conn3, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection\n")
	defer conn3.Close()

	// Creamos conexion con el servicio 
	csp := lj.NewLiderJugadorServiceClient(conn3)
	// Conectamos con el lider y se imprime la respuesta
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	respuesta,err2 := csp.EtapaUno(ctx, &lj.NumPasosReq{PlayMsg: numero, NroJugador: nroJugadorAux, Total: contador, Ronda: ronda})
	ut.FailOnError(err2, "Failed to send a play\n")

	return respuesta.GetStateMsg()
}

// Juego 1 por rondas (4 rondas maximo)
func StartGameUno(wg *sync.WaitGroup){
	defer wg.Done()
	var eleccion int64
	nroAux := rg.GetNumeroJugador()
	vivosAux := rg.GetVivosSlice()
	fmt.Println("\nInicio Juego 1")
	fmt.Println("Debe ingresar numeros entre 1 y 10 durante 4 rondas")
	fmt.Println("Usted muere si no suma 21 al final de la cuarta ronda")
	fmt.Println("Usted tambien muere si supera a la eleccion del lider")
	fmt.Println("Mucha suerte")

	var sumador int64 = 0
	for i:=0; i<4; i++{
		eleccion = 0
		fmt.Println("\nInicio Ronda", i+1)

		fmt.Println("Ingrese un numero entre 1 y 10:")
		fmt.Scanln(&eleccion)
		time.Sleep(3*time.Second)

		for (eleccion>10 || eleccion<1){
			fmt.Println("Vuelva a ingresar el numero. Debe estar entre 1 y 10")
			fmt.Scanln(&eleccion)
		}

		aux := eleccion
		resp := EtapaUno(aux, sumador, int64(i), nroAux)
		if resp==0{
			fmt.Println("Tu estas muerto, BANG!!")
			fmt.Println("Tu numero fue", aux)
			vivosAux[nroAux-1] = false
			break
		}
		sumador += aux
		if (sumador>=21 || resp==3){
			fmt.Println("Superaste el juego 1, Felicidades!!")
			fmt.Println("Tu numero fue", aux)
			fmt.Println("Tu total fue", sumador)
			break
		}
		if i==3{
			if (sumador<21 || resp==2){
				fmt.Println("No lograste llegar al final, tu estas muerto, BANG!!")
				fmt.Println("Tu numero fue", aux)
				fmt.Println("Tu total fue", sumador)
				vivosAux[nroAux-1] = false
				break
			} else if (sumador>=21 || resp==3){
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

func StartGameUnoBot(numeroJugadorBot int64, wg *sync.WaitGroup){
	defer wg.Done()
	var sumador int64 = 0
	vivosAux := rg.GetVivosSlice()
	for i:=0; i<4; i++{
		aux := RandomNumber(int64(i)*numeroJugadorBot)
		resp := EtapaUno(aux, sumador, int64(i), numeroJugadorBot)
		if resp==0{
			vivosAux[numeroJugadorBot-1] = false
			break
		}
		sumador += aux
		if (sumador>=21 || resp==3){
			break
		}
		
		if i==3{
			if (sumador<21 || resp==2){
				vivosAux[numeroJugadorBot-1] = false
				break
			}
		}
		time.Sleep(10*time.Second)
	}
}

func StartGameUnoTrigger(){
	// wait groups
	var wg sync.WaitGroup
	wg.Add(rg.GetMaxJug())
	for i:=0; i<rg.GetMaxJug(); i++{
		if i!=0{
			go StartGameUnoBot(int64(i+1), &wg)
		} else {
			go StartGameUno(&wg)
		}
	}

	wg.Wait()
}