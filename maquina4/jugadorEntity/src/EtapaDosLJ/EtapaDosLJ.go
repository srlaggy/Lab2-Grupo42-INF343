package EtapaDosLJ

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
	port_grpc = "50002"
	min = 1
	max = 4
)

//var global
var vivosInicioDos []int64
var vivosFinDos []bool
var conectarSlice []bool
// var jugaronSlice []bool
var savePos int64

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: Etapa 2 ----- // --> Jugador actua como cliente

func Etapa2Conn(nroAux int64) {
	// Creamos conexion
	conn3, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn3.Close()

	// Creamos conexion con el servicio 
	csp := lj.NewLiderJugadorServiceClient(conn3)
	// Conectamos con el lider y se imprime la respuesta
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r,err2 := csp.Etapa2Conn(ctx, &lj.E2ConnReq{NroJugador: nroAux})
	ut.FailOnError(err2, "Failed to send a petition")
	if (nroAux==1){
		if (r.GetNroGroup() == 0){
			fmt.Println("Mala suerte, quedaste fuera. BANG!!")
			vivosFinDos[FindIndex(nroAux, vivosInicioDos)] = false
		} else {
			fmt.Println("Fuiste asignado al grupo", r.GetNroGroup(), "para el siguiente juego")
		}
	} else if (r.GetNroGroup() == 0){
		vivosFinDos[FindIndex(nroAux, vivosInicioDos)] = false
	}
	conectarSlice[FindIndex(nroAux, vivosInicioDos)] = true
}

// inicializa conexion de todos a etapa 2
func Etapa2ConnTrigger(){
	vivosInicioDos = GetVivosInicioDos()
	time.Sleep(2*time.Second)
	conectarSlice = make([]bool, len(vivosInicioDos))
	vivosFinDos = make([]bool, len(vivosInicioDos))
	for i:=0; i<len(vivosInicioDos); i++{
		conectarSlice[i] = false
		vivosFinDos[i] = true
		if vivosInicioDos[i]!=1{
			go Etapa2Conn(vivosInicioDos[i])
		} else {
			savePos = int64(i)
		}
	}
	Etapa2Conn(vivosInicioDos[savePos])

	for{
		if RevConectaron(){
			break
		}
	}
}

func FindIndex(num int64, lista []int64) int{
	for i:=0; i<len(lista); i++{
		if lista[i]==num{
			return i
		}
	}
	return -1
}

// funcion que revisa si todos se conectaron a etapa 2
func RevConectaron() bool{
	// retorna si todos conectaron
	conectar := true
	for i:=0; i<len(conectarSlice); i++{
		// si pillo uno que no haya jugado, en verdad no han jugado todos
		if !conectarSlice[i]{
			conectar = false
		}
	}
	// true -> todos conectar
	// false -> NO todos conectar
	return conectar
}

// detectar cuantos vivos hay
func GetVivosInicioDos() []int64{
	aux := make([]int64, 0, rg.GetMaxJug())
	aux2 := rg.GetVivosSlice()
	aux3 := rg.GetNumeroJugadorBot()
	for i:=0; i<rg.GetMaxJug(); i++{
		if aux2[i]{
			aux = append(aux, aux3[i])
		}
	}
	return aux
}

func GetVivosFinDos() []bool{
	return vivosFinDos
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