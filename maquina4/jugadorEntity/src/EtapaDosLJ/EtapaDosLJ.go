package EtapaDosLJ

import (
	"context"
	"time"
	"sync"
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
	port2 = "50100"
	min = 1
	max = 4
)

//var global

// son los numeros vivos al comienzo de la etapa 2 - preasignacion de grupos
var vivosInicioDos []int64
// son los vivos descontando el posible personaje sin grupo (son booleanos)
var vivosFinDos []bool
// indice de la posicion del jugador humano dentro del slice de jugadores vivos (vivosInicioDos)
var savePos int64
// slice de vivos luego de formacion de grupos y su correspondiente vivo o muerto
var vivosNumeros []int64
var vivosBool []bool

var grupos []int64

//final segundo juego
var vivosNumerosFinal []int64
var vivosBoolFinal []bool

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: Etapa 2 ----- // --> Jugador actua como cliente

//------------------------------------------------------//
//------------------CONEXIONES--------------------------//
//------------------------------------------------------//

func Etapa2Conn(nroAux int64, wg *sync.WaitGroup) {
	// terminar routine en wait group
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
	r,err2 := csp.Etapa2Conn(ctx, &lj.E2ConnReq{NroJugador: nroAux})
	ut.FailOnError(err2, "Failed to send a petition")
	if (nroAux==1){
		if (r.GetNroGroup() == 0){
			fmt.Println("\nMala suerte, quedaste fuera. BANG!!\n")
			vivosFinDos[FindIndex(nroAux, vivosInicioDos)] = false
		} else {
			fmt.Println("\nFuiste asignado al grupo", r.GetNroGroup(), "para el siguiente juego\n")
		}
	} else if (r.GetNroGroup() == 0){
		vivosFinDos[FindIndex(nroAux, vivosInicioDos)] = false
	}
	grupos[FindIndex(nroAux, vivosInicioDos)] = r.GetNroGroup()
}

// inicio de etapa 2 - conexion con request
func Etapa2(nroEleccion int64, nroJugador int64) int64{
	// Creamos conexion
	conn3, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn3.Close()

	// Creamos conexion con el servicio 
	csp := lj.NewLiderJugadorServiceClient(conn3)
	// Conectamos con el lider y se imprime la respuesta
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r,err2 := csp.Etapa2(ctx, &lj.Etapa2Req{NroJugador: nroJugador, Numero: nroEleccion})
	ut.FailOnError(err2, "Failed to send a response")
	return r.GetStateMsg()
}

func Etapa2Final(nroJugadorAux int64, nroGrupoAux int64) int64{
	// Creamos conexion
	conn4, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port2), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn4.Close()

	// Creamos conexion con el servicio 
	csp := lj.NewLiderJugadorServiceClient(conn4)
	// Conectamos con el lider y se imprime la respuesta
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r,err2 := csp.Etapa2Fin(ctx, &lj.E2FinReq{NroJugador: nroJugadorAux, NroGroup: nroGrupoAux})
	ut.FailOnError(err2, "Failed to send a petition")
	return r.GetDead()
}

//------------------------------------------------------//
//--------------------TRIGGERS--------------------------//
//------------------------------------------------------//

// inicializa conexion de todos a etapa 2
func Etapa2ConnTrigger(){
	var auxWait int
	vivosInicioDos = GetVivosInicioDos()
	grupos = make([]int64, len(vivosInicioDos))
	vivosFinDos = make([]bool, len(vivosInicioDos))
	// largo del wait group
	if (len(vivosInicioDos)%2 != 0){
		auxWait = len(vivosInicioDos) - 1
	} else {
		auxWait = len(vivosInicioDos)
	}
	time.Sleep(5*time.Second)
	// wait group
	var wg sync.WaitGroup
	wg.Add(auxWait)
	time.Sleep(5*time.Second)
	for i:=0; i<len(vivosInicioDos); i++{
		vivosFinDos[i] = true
		if vivosInicioDos[i]!=1{
			go Etapa2Conn(vivosInicioDos[i], &wg)
		} else {
			savePos = int64(i)
			go Etapa2Conn(vivosInicioDos[savePos], &wg)
		}
	}
	wg.Wait()
}

//------------------------------------------------------//
//--------------------JUGADORES-------------------------//
//------------------------------------------------------//

// trigger de jugador humano para request
func HumanoEtapa2(indice int64, wg2 *sync.WaitGroup){
	defer wg2.Done()
	var eleccion int64
	fmt.Println("\nInicio Juego 2")
	fmt.Println("Debe ingresar un numero entre 1 y 4")
	fmt.Println("Usted muere si la paridad de la suma de su equipo es distinta a la eleccion del lider")
	fmt.Println("Buena suerte")

	eleccion = 0
	fmt.Println("\nInicio Ronda 1")

	fmt.Println("Ingrese un numero entre 1 y 4:")
	fmt.Scanln(&eleccion)
	time.Sleep(3*time.Second)

	for (eleccion>4 || eleccion<1){
		fmt.Println("Vuelva a ingresar el numero. Debe estar entre 1 y 4")
		fmt.Scanln(&eleccion)
	}

	aux := eleccion
	noUsar := Etapa2(aux, vivosNumeros[indice])
	_ = noUsar
	resp := Etapa2Final(vivosNumeros[indice], grupos[FindIndex(vivosNumeros[indice], vivosInicioDos)])
	// 1 muere
	// 0 vive
	if resp==1{
		fmt.Println("\nTu equipo es muy debil, BANG!!\n")
		vivosBool[indice] = false
	} else if resp==0{
		fmt.Println("\nGran equipo, pasan a la siguiente Etapa\n")
	}
	time.Sleep(3*time.Second)
}

// trigger de jugador bot para request
func BotEtapa2(indice int64, wg2 *sync.WaitGroup){
	defer wg2.Done()
	aux := RandomNumber(indice*5)
	noUsar := Etapa2(aux, vivosNumeros[indice])
	_ = noUsar
	resp := Etapa2Final(vivosNumeros[indice], grupos[FindIndex(vivosNumeros[indice], vivosInicioDos)])
	// 1 muere
	// 0 vive
	if resp==1{
		vivosBool[indice] = false
	}
	time.Sleep(3*time.Second)
}

//------------------------------------------------------//
//------------------LOOP-E-INICIO-----------------------//
//------------------------------------------------------//

// trigger general etapa 2 (juego 2)
func StartEtapa2Trigger(){
	vivosNumeros, vivosBool = CreateVivosAndBool()
	// numero de vivos con grupos creados
	var numAuxx int = NumVivos()
	// wait group
	var wg2 sync.WaitGroup
	wg2.Add(numAuxx)
	for i:=0; i<numAuxx; i++{
		if vivosNumeros[i]!=1{
			// indice en arreglo y numero de jugador
			go BotEtapa2(int64(i), &wg2)
		} else {
			go HumanoEtapa2(savePos, &wg2)
		}
	}

	wg2.Wait()

	vivosNumerosFinal, vivosBoolFinal = setFinalValues()
}

// detectar cuantos vivos hay y entregarlos en un slice
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

func CreateVivosAndBool() ([]int64, []bool){
	var aux int = NumVivos()
	jugadoresAux := make([]int64, 0, aux)
	vivosAux := make([]bool, 0, aux)
	for i:=0; i<len(vivosFinDos); i++{
		if vivosFinDos[i]{
			jugadoresAux = append(jugadoresAux, vivosInicioDos[i])
			vivosAux = append(vivosAux, true)
		}
	}
	return jugadoresAux, vivosAux
}

//------------------------------------------------------//
//----------------------UTILES--------------------------//
//------------------------------------------------------//

// funcion que permite saber cuantos vivos hay luego de la reparticion de grupos
func NumVivos() int{
	var aux int = 0
	for i:=0; i<len(vivosFinDos); i++{
		if vivosFinDos[i]{
			aux += 1
		}
	}
	return aux
}

// determina numero random del jugador
func RandomNumber(num int64) int64{
	rand.Seed(time.Now().UnixNano() + num*3)
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

// funcion que setea vivos finales
func setFinalValues() ([]int64, []bool){
	aux1 := make([]int64, 0, len(vivosNumeros))
	aux2 := make([]bool, 0, len(vivosNumeros))
	for i:=0; i<len(vivosNumeros); i++{
		if vivosBool[i]{
			aux1 = append(aux1, vivosNumeros[i])
			aux2 = append(aux2, vivosBool[i])
		}
	}
	return aux1, aux2
}

//------------------------------------------------------//
//----------------------GETTERS-------------------------//
//------------------------------------------------------//
func GetVivosNumeros() []int64{
	return vivosNumeros
}

func GetVivosBool() []bool{
	return vivosBool
}

func GetVivosNumerosFinal() []int64{
	return vivosNumerosFinal
}

func GetVivosBoolFinal() []bool{
	return vivosBoolFinal
}