package EtapaTresLJ

import (
	"context"
	"time"
	"fmt"
	"math/rand"
	"sync"
	"google.golang.org/grpc"
	ut "lab/jugador/utils"
	lj "lab/jugador/proto/LJ"
	// rg "lab/jugador/src/requestGameLJ"
	e2 "lab/jugador/src/EtapaDosLJ"
)

const (
	address = "localhost"
	protocolo_grpc = ""
	port_grpc = "50003"
	min = 1
	max = 10
)

//var global
var nromuerto int = 0
var jugadoresE3 []int64
var vivosE3 []bool
var cantVivos int
var numGrp []int64
var respuesta []int64

//------------------------------------------------------//
//------------------CONEXIONES--------------------------//
//------------------------------------------------------//

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: Etapa 2 ----- // --> Jugador actua como cliente

func Etapa3Conn(nroJugador int64, wg *sync.WaitGroup) {
	defer wg.Done()
	// Creamos conexion
	conn3, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn3.Close()

	// Creamos conexion con el servicio 
	csp := lj.NewLiderJugadorServiceClient(conn3)
	// Conectamos con el lider y se imprime la respuesta
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r,err2 := csp.Etapa2Conn(ctx, &lj.E2ConnReq{NroJugador: nroJugador})
	ut.FailOnError(err2, "Failed to send a petition")
	numGrp[FindIndex(nroJugador, jugadoresE3)] = r.GetNroGroup()
	if (int(r.GetNroGroup()) != 10){
		fmt.Println("Mi numero de grupo es:", r.GetNroGroup())
	}else{
		fmt.Println("Mala suerte, he muerto porque el total de jugadores era impar", nroJugador)
		nromuerto = int(nroJugador)
		vivosE3[FindIndex(nroJugador, jugadoresE3)] = false
		cantVivos -= 1
	}
}

func Etapa3(nroJugador int64, numero int64) {
	// Creamos conexion
	fmt.Println("Entre")
	conn3, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn3.Close()

	// Creamos conexion con el servicio 
	csp := lj.NewLiderJugadorServiceClient(conn3)
	// Conectamos con el lider y se imprime la respuesta
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r,err2 := csp.Etapa2(ctx, &lj.Etapa2Req{NroJugador: nroJugador, Numero: numero})
	ut.FailOnError(err2, "Failed to send a play")
	respuesta[FindIndex(nroJugador, jugadoresE3)] = r.GetStateMsg()
	fmt.Println(r.GetStateMsg())
}

//------------------------------------------------------//
//--------------------TRIGGERS--------------------------//
//------------------------------------------------------//

func Etapa3ConnTrigger(){
	jugadoresE3 = e2.GetVivosNumerosFinal()
	vivosE3 = e2.GetVivosBoolFinal()
	cantVivos = len(jugadoresE3)
	numGrp = make([]int64, cantVivos)
	respuesta = make([]int64, cantVivos)
	time.Sleep(5*time.Second)

	var wg sync.WaitGroup
	wg.Add(cantVivos)
	for i:=0; i<cantVivos; i++{
		go Etapa3Conn(jugadoresE3[i], &wg)
	}
	wg.Wait()
}

func Etapa3Trigger(){
	var wg2 sync.WaitGroup
	wg2.Add(cantVivos)
	for i:=0; i<cantVivos; i++{
		if jugadoresE3[i]==1{
			go HumanE3(jugadoresE3[i], &wg2)
		} else {
			// go Etapa3Conn(jugadoresE3[i], &wg)
			go BotE3(jugadoresE3[i], &wg2)
		}
	}
	wg2.Wait()
}

//------------------------------------------------------//
//--------------------JUGADORES-------------------------//
//------------------------------------------------------//

func HumanE3(nroJugador int64, wg *sync.WaitGroup){
	defer wg.Done()

	var eleccion int64
	fmt.Println("\nInicio Juego 3")
	fmt.Println("Debe ingresar un numero entre 1 y 10")
	fmt.Println("Usted muere si el calculo del lider no le favorece")
	fmt.Println("Buena suerte, estas llegando al final!!!")

	eleccion = 0
	fmt.Println("\nInicio Ronda 1")

	fmt.Println("Ingrese un numero entre 1 y 10:")
	fmt.Scanln(&eleccion)
	time.Sleep(3*time.Second)

	for (eleccion>10 || eleccion<1){
		fmt.Println("Vuelva a ingresar el numero. Debe estar entre 1 y 10")
		fmt.Scanln(&eleccion)
	}

	aux := eleccion
	Etapa3(nroJugador, int64(aux))
	time.Sleep(5*time.Second)
	// 1 vive
	// 0 muere
	if respuesta[FindIndex(nroJugador, jugadoresE3)]==0{
		fmt.Println("\nCasi lo logras, BANG!!\n")
		vivosE3[FindIndex(nroJugador, jugadoresE3)] = false
		cantVivos -= 1
	} else if respuesta[FindIndex(nroJugador, jugadoresE3)]==1{
		fmt.Println("\nImpresionante!! Eres el ganador de lso jeugos del Calamar\n")
	}
	time.Sleep(5*time.Second)
}

func BotE3(nroJugador int64, wg *sync.WaitGroup){
	defer wg.Done()

	aux := RandomNumber(nroJugador)
	Etapa3(nroJugador, aux)
	time.Sleep(5*time.Second)
	// 1 vive
	// 0 muere
	if respuesta[FindIndex(nroJugador, jugadoresE3)]==0{
		vivosE3[FindIndex(nroJugador, jugadoresE3)] = false
		cantVivos -= 1
	}
	time.Sleep(5*time.Second)
}

//------------------------------------------------------//
//----------------------UTILES--------------------------//
//------------------------------------------------------//

// determina numero random del jugador
func RandomNumber(num int64) int64{
	rand.Seed(time.Now().UnixNano() + num*15)
	aux := rand.Intn(max - min) + min
	return int64(aux)
}

// funcion que busca indice de elemento en una lista
func FindIndex(num int64, lista []int64) int{
	for i:=0; i<len(lista); i++{
		if lista[i]==num{
			return i
		}
	}
	return -1
}

//------------------------------------------------------//
//----------------------GETTERS-------------------------//
//------------------------------------------------------//

func GetNroMuerto() int{
	return nromuerto
}

func GetVivosE3() []bool{
	return vivosE3
}

func GetJugadoresE3() []int64{
	return jugadoresE3
}

func GetCantVivos() int{
	return cantVivos
}