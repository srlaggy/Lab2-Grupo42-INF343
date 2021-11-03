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

// variables globales
var jugaronSlice []bool

func RandomNumber() int64{
	rand.Seed(time.Now().UnixNano())
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
func StartGameUno(){
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
			jugaronSlice[nroAux-1] = true
			vivosAux[nroAux-1] = false
			break
		}
		sumador += aux
		if (sumador>=21 || resp==3){
			fmt.Println("Superaste el juego 1, Felicidades!!")
			fmt.Println("Tu numero fue", aux)
			fmt.Println("Tu total fue", sumador)
			jugaronSlice[nroAux-1] = true
			break
		}
		if i==3{
			if (sumador<21 || resp==2){
				fmt.Println("No lograste llegar al final, tu estas muerto, BANG!!")
				fmt.Println("Tu numero fue", aux)
				fmt.Println("Tu total fue", sumador)
				jugaronSlice[nroAux-1] = true
				vivosAux[nroAux-1] = false
				break
			} else if (sumador>=21 || resp==3){
				fmt.Println("Superaste el juego 1, Felicidades!!")
				fmt.Println("Tu numero fue", aux)
				fmt.Println("Tu total fue", sumador)
				jugaronSlice[nroAux-1] = true
				break
			}
		}
		fmt.Println("Te salvaste esta vez")
		fmt.Println("Tu numero fue", aux)
		fmt.Println("Tu total acumulado es de", sumador)
		time.Sleep(3*time.Second)
	}
}

func StartGameUnoBot(numeroJugadorBot int64){
	var sumador int64 = 0
	vivosAux := rg.GetVivosSlice()
	for i:=0; i<4; i++{
		aux := RandomNumber()
		// fmt.Println("El numero del jugador", numeroJugadorBot, "es", aux, "en ronda", i+1)
		resp := EtapaUno(aux, sumador, int64(i), numeroJugadorBot)
		if resp==0{
			jugaronSlice[numeroJugadorBot-1] = true
			vivosAux[numeroJugadorBot-1] = false
			break
		}
		sumador += aux
		if (sumador>=21 || resp==3){
			jugaronSlice[numeroJugadorBot-1] = true
			break
		}
		if i==3{
			if (sumador<21 || resp==2){
				jugaronSlice[numeroJugadorBot-1] = true
				vivosAux[numeroJugadorBot-1] = false
				break
			} else {
				jugaronSlice[numeroJugadorBot-1] = true
				break
			}
		}
		time.Sleep(10*time.Second)
	}
}

func StartGameUnoTrigger(){
	jugaronSlice = make([]bool, rg.GetMaxJug())
	// inicializo valores de jugado - cuando todos hayan jugando puede cerrarse la conexion de juego 1
	for i:=0; i<rg.GetMaxJug(); i++{
		jugaronSlice[i] = false
		if i!=0{
			go StartGameUnoBot(int64(i+1))
		}
	}
	StartGameUno()

	for{
		if RevJugaron(){
			break
		}
	}
}

// funcion que revisa si todos jugaron
func RevJugaron() bool{
	// retorna si todos jugaron
	jugaron := true
	for i:=0; i<len(jugaronSlice); i++{
		// si pillo uno que no haya jugado, en verdad no han jugado todos
		if !jugaronSlice[i]{
			jugaron = false
		}
	}
	// true -> todos jugaron
	// false -> NO todos jugaron
	return jugaron
}