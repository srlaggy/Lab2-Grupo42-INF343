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

// son los numeros vivos al comienzo de la etapa 2 - preasignacion de grupos
var vivosInicioDos []int64
// son los vivos descontando el posible personaje sin grupo (son booleanos)
var vivosFinDos []bool
// son un slice de control que permite saber cuando todos hayan enviado su peticion de conexion al juego
var conectarSlice []bool
// slice de control que permite saber quienes ya jugaron
var jugaronSlice []bool
// indice de la posicion del jugador humano dentro del slice de jugadores vivos (vivosInicioDos)
var savePos int64
// slice de vivos luego de formacion de grupos y su correspondiente vivo o muerto
var vivosNumeros []int64
var vivosBool []bool

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

// funcion que busca indice de elemento en una lista
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

// determina numero random del jugador
func RandomNumber() int64{
	rand.Seed(time.Now().UnixNano())
	aux := rand.Intn(max - min) + min
	return int64(aux)
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

// trigger de jugador humano para request
func HumanoEtapa2(indice int64){
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
	resp := Etapa2(aux, vivosNumeros[indice])
	// 0 muere
	// 1 vive
	if resp==0{
		fmt.Println("Tu equipo es muy debil, BANG!!")
		vivosBool[indice] = false
		jugaronSlice[indice] = true
	} else if resp==1{
		fmt.Println("Gran equipo, pasan a la siguiente Etapa")
		jugaronSlice[indice] = true
	}
	time.Sleep(3*time.Second)
}

// trigger de jugador bot para request
func BotEtapa2(indice int64){
	aux := RandomNumber()
	resp := Etapa2(aux, vivosNumeros[indice])
	// 0 muere
	// 1 vive
	if resp==0{
		vivosBool[indice] = false
		jugaronSlice[indice] = true
	} else if resp==1{
		jugaronSlice[indice] = true
	}
	time.Sleep(3*time.Second)
}

func CreateVivosAndBool() ([]int64, []bool, []bool){
	var aux int = NumVivos()
	jugadoresAux := make([]int64, 0, aux)
	vivosAux := make([]bool, 0, aux)
	jugaronAux := make([]bool, 0, aux)
	for i:=0; i<len(vivosFinDos); i++{
		if vivosFinDos[i]{
			jugadoresAux = append(jugadoresAux, vivosInicioDos[i])
			vivosAux = append(vivosAux, true)
			jugaronAux = append(jugaronAux, false)
		}
	}
	return jugadoresAux, vivosAux, jugaronAux
}

// trigger general etapa 2 (juego 2)
func StartEtapa2Trigger(){
	vivosNumeros, vivosBool, jugaronSlice = CreateVivosAndBool()
	// inicializo valores de jugado - cuando todos hayan jugando puede cerrarse la conexion de juego 1
	for i:=0; i<NumVivos(); i++{
		// si es humano no entra
		if vivosNumeros[i]!=1{
			// indice en arreglo y numero de jugador
			go BotEtapa2(int64(i))
		}
	}
	HumanoEtapa2(savePos)

	for{
		if RevJugaron(){
			break
		}
	}
}

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

// funcion que revisa si todos jugaron la etapa 2
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

func GetVivosNumeros() []int64{
	return vivosNumeros
}

func GetVivosBool() []bool{
	return vivosBool
}