package EtapaTresLJ

import (
	"context"
	// "log"
	"net"
	"fmt"
	"math/rand"
	"math"
	"time"
	"google.golang.org/grpc"
	lj "lab/lider/proto/LJ"
	sg "lab/lider/src/startGameL"
	ut "lab/lider/utils"
)

const (
	address = "localhost"
	protocolo_grpc = "tcp"
	port_grpc1 = "50003"
	min = 1
	max = 10
)

// var global
var lista_parejas [][]int64
var lista_num_parejas [][]int64
var jugadores_inter_etapa3 int = 0
var nroLider int64
var sgv []bool
var num_eliminado int64 = 0
var vivostotales int

// ----- FUNCIÓN: recibir jugadas del jugador etapa 3 ----- // --> Lider actua como servidor
type server struct {
	lj.UnimplementedLiderJugadorServiceServer
}

//------------------------------------------------------//
//------------------CONEXIONES--------------------------//
//------------------------------------------------------//

// funcion: tercer juego (reutilizamos protos y servicios)
func (s *server) Etapa2Conn(ctx context.Context, in *lj.E2ConnReq) (*lj.E2ConnResp, error) {
	fmt.Println("Ingreso el jugador: ", in.NroJugador)
	if (in.NroJugador == num_eliminado){
		return &lj.E2ConnResp{NroGroup: 10}, nil
	}
	nro_group_jugador := contains(lista_parejas, in.NroJugador)
	return &lj.E2ConnResp{NroGroup: int64(nro_group_jugador)}, nil // tu moriste
}

func (s *server) Etapa2(ctx context.Context, in *lj.Etapa2Req) (*lj.Etapa2Resp, error) {
	fmt.Println("Entra a jugar: ", in.NroJugador)
	nro_group_jugador_i := contains(lista_parejas, in.NroJugador)
	nro_group_jugador_j := contains_subslice(lista_parejas, nro_group_jugador_i, in.NroJugador)
	lista_num_parejas[nro_group_jugador_i][nro_group_jugador_j] = in.Numero
	// loop
	jugadores_inter_etapa3 += 1
	fmt.Println(in.NroJugador, "esperando y su numero fue:", in.Numero, ", jugadores en la etapa:", jugadores_inter_etapa3)
	for (jugadores_inter_etapa3 < vivostotales){
	}
	valor_bool := comparacion_final(nro_group_jugador_i, nro_group_jugador_j)
	fmt.Println("jugador:", in.NroJugador, "se salva?", valor_bool)
	// 1 - vive
	// 0 - muere
	if (valor_bool == true){
		return &lj.Etapa2Resp{StateMsg: int64(1)}, nil
	}else{
		return &lj.Etapa2Resp{StateMsg: int64(0)}, nil
	}
}

//------------------------------------------------------//
//----------------------LOGICA--------------------------//
//------------------------------------------------------//

// funcion que compara los numeros de una pareja en la posicion i de la lista de numeros
func comparacion_final(pos_i int, pos_j int) bool{
	num_jugador1 := lista_num_parejas[pos_i][0]
	num_jugador2 := lista_num_parejas[pos_i][1]
	resta1 := math.Abs(float64(nroLider-num_jugador1))
	resta2 := math.Abs(float64(nroLider-num_jugador2))
	list := make([]bool , 2)
	if (resta1 < resta2){
		list[0] = true
		list[1] = false
	}else if (resta1 > resta2){
		list[0] = false
		list[1] = true
	}else{
		list[0] = true
		list[1] = true
	}
	return list[pos_j]
}

// funcion elimina un jugador en caso que el numero total de jugadores sea impar, sino, lo mantiene
// True: cantidad par
// False: cantidad impar
func ParidadEtapa3() (int, bool, []int64){
	c := 0
	sgj := sg.GetJugadores()
	jugadores_vivos := make([]int64, 0, sg.GetCantidadJugadores()) 		// se crea una lista para guardar el id de los numeros vivos
	for j := 0; j < len(sgv); j++ {
		if (sgv[j] == true){
			jugadores_vivos = append(jugadores_vivos, int64(sgj[j]))
			c++
		}
	}
	fmt.Println("Jugadores vivos:", jugadores_vivos)
	if(c % 2 == 0){
		return c, true, jugadores_vivos
	}else{
		return c, false, jugadores_vivos
	}
}

// funcion: hacer grupos
func Parejas(){
	rand.Seed(time.Now().UnixNano())
	cantvivos, paridad, lista_vivos := ParidadEtapa3()
	// si la cantidad de numeros vivos es impar entonces se debe eliminar un jugador aleatoriamente
	if (!paridad){
		nro_azar := int64(rand.Intn(cantvivos-1))
		num_eliminado = lista_vivos[nro_azar]			// el idJugador
		sg.SetVivos(int(num_eliminado)-1, false)		// se setea a false el jugador
		cantvivos = cantvivos-1
		fmt.Println("Nro eliminado:", num_eliminado, ", cantidad de vivos:", cantvivos)
		fmt.Println("lista_vivos antes de remove:", lista_vivos)
		lista_vivos := remove(lista_vivos, FindIndex(num_eliminado, lista_vivos))
		fmt.Println("lista_vivos luego de remove:", lista_vivos)
	}
	vivostotales = cantvivos
	// se crea una lista aleatoria entre 1 y  la cantidad de jugadores vivos cantvivos
	list := rand.Perm(cantvivos)
	for i, _ := range list {
		list[i]++
	}
	// se crea una lista de lista de parejas
	lista_parejas = make([][]int64 , cantvivos/2)
	lista_num_parejas = make([][]int64 , cantvivos/2)
	count := 0
	
	for i := 0; i < cantvivos/2; i++ {
        // se declaran los slices de slices de tamaño 2 (parejas)
        lista_parejas[i] = make([]int64, 2)
        lista_num_parejas[i] = make([]int64, 2)
        // se definiran las parejas según la list
        for j := 0; j < 2; j++{
			lista_parejas[i][j] = lista_vivos[list[count]-1] 	// con list[count]: se obtiene el numero random-1 y luego con eso se obtiene el numero del jugador
			lista_num_parejas[i][j] = int64(0)					// se rellena con 0 la lista con numeros de los jugadores
			count ++
        }
    }
}

//------------------------------------------------------//
//----------------------UTILES--------------------------//
//------------------------------------------------------//

// funcion: encontrar la indexación de un elemento de un slice
func FindIndex(num int64, lista []int64) int{
    for i:=0; i<len(lista); i++{
        if lista[i]==num{
            return i
        }
    }
    return -1
}

// funcion: eliminar un elemento en específico de un slice
func remove(slice []int64, s int) []int64 {
    return append(slice[:s], slice[s+1:]...)
}

// funcion: lider elige su numero aleatoriamente enter 1 y 10
func NroLider(){
	rand.Seed(time.Now().UnixNano())
	nroLider = int64(rand.Intn(max-min) + min)
}

// funcion: contains
// retorna indice de la pareja en lista de parejas
func contains(s [][]int64, num int64) int {
	count := 0
    for _, v := range s {
        if v[0] == num {
            return count
        }else if v[1] == num {
            return count
        }
		count++
    }
    return 0
}

// funcion: contains
// retorna si es la primera (0) o segunda (1) persona de la pareja
func contains_subslice(s [][]int64, pos_i int, num int64) int {
    for _, v := range s {
        if v[0] == num {
            return 0
        }else if v[1] == num {
            return 1
        }
    }
    return 5
}

// setea variables iniciales antes de iniciar servidor
func SetterEtapa(){
	sgv = sg.GetVivos()
	Parejas()
	NroLider()
}

//------------------------------------------------------//
//----------------------GETTERS-------------------------//
//------------------------------------------------------//

// funcion: obtener el numero del lider
func GetNroLider() int64{
	return nroLider
}

//------------------------------------------------------//
//----------------------REQUEST-------------------------//
//------------------------------------------------------//

// funciones: crea la conexión
func Grpc_func() {
	fmt.Println("Lista parejas inicio:", lista_parejas)
	fmt.Println("El numero del lider es:", nroLider)
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc1)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	lj.RegisterLiderJugadorServiceServer(s, &server{})
	// log.Printf("Servidor grpc escuchando en el puerto %v", port_grpc1)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}

///////////////////////////////////////////////////////////////////////////////////////////////
func LoopAux(){
	for{
		if jugadores_inter_etapa3 >= vivostotales{
			break
		}
	}
	fmt.Println("60 segundos para el cierre")
	time.Sleep(60*time.Second)
}