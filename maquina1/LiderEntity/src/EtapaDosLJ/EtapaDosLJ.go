package EtapaDosLJ

import (
	"context"
	"log"
	"net"
	"fmt"
	"math/rand"
	"time"
	"google.golang.org/grpc"
	lj "lab/lider/proto/LJ"
	sg "lab/lider/src/startGameL"
	ut "lab/lider/utils"
)

const (
	address = "localhost"
	protocolo_grpc = "tcp"
	port_grpc1 = "50002"
	min = 1
	max = 4
)

// var global
var num_group1 int64 = 0
var num_group2 int64 = 0
var group1 []int64
var group2 []int64
var jugadores_inter_etapa2 int = 0
var nroLider int64

// ----- FUNCIÓN: recibir jugadas del jugador ----- // --> Lider actua como servidor
type server struct {
	lj.UnimplementedLiderJugadorServiceServer
}

// funcion: segundo juego
// muerto 0 - vivo 1
func (s *server) Etapa2Conn(ctx context.Context, in *lj.E2ConnReq) (*lj.E2ConnResp, error) {
	if (contains(group1, in.NroJugador)){
		return &lj.E2ConnResp{NroGroup: int64(1)}, nil // grupo 1
	}
	if (contains(group2, in.NroJugador)){
		return &lj.E2ConnResp{NroGroup: int64(2)}, nil // grupo 2
	}else{
		// vivo = false
		sg.GetVivos()[in.NroJugador-1] = false
		return &lj.E2ConnResp{NroGroup: int64(0)}, nil // tu moriste
	}
}

func (s *server) Etapa2(ctx context.Context, in *lj.Etapa2Req) (*lj.Etapa2Resp, error) {
	jugadores_inter_etapa2 += 1
	if (contains(group1, in.NroJugador)){
		num_group1 = num_group1 + in.Numero
	}else if (contains(group2, in.NroJugador)){
		num_group2 = num_group2 + in.Numero
	}
	fmt.Println(in.NroJugador, "esperando y su numero fue:", in.Numero)
	// loop
	for (jugadores_inter_etapa2 < len(group1)+len(group2)){
	}
	// return
	if (contains(group1, in.NroJugador)){
		return &lj.Etapa2Resp{StateMsg: int64(num_group1)}, nil // tu moriste
	}else{
		return &lj.Etapa2Resp{StateMsg: int64(num_group2)}, nil // tu moriste
	}
}

// funcion: contains
func contains(s []int64, num int64) bool {
    for _, v := range s {
        if v == num {
            return true
        }
    }
    return false
}

// funcion: hacer grupos
func GroupAndNumberLider(){
	sgj := sg.GetJugadores()
	sgv := sg.GetVivos()
	jugadores_vivos := make([]int64, 0, sg.GetCantidadJugadores())
	rand.Seed(time.Now().UnixNano())
	// Obtenemos los numeros de los jugadores vivos
	for j := 0; j < len(sgj); j++ {
		if (sgv[j] == true){
			jugadores_vivos = append(jugadores_vivos, int64(sgj[j]))
		}
	}
	// se declaran las listas de grupos
	group1 = make([]int64, 0, len(jugadores_vivos)/2)
	group2 = make([]int64, 0, len(jugadores_vivos)/2)

	list := rand.Perm(len(jugadores_vivos))
	for i, _ := range list {
		list[i]++
	}
	for i := 0; i < len(jugadores_vivos); i++ {
		if (len(group1)<len(jugadores_vivos)/2){
			group1 = append(group1, jugadores_vivos[list[i]-1])
		} else if (len(group2)<len(jugadores_vivos)/2){
			group2 = append(group2, jugadores_vivos[list[i]-1])
		}
	}

	nroLider = int64(rand.Intn(max-min) + min)
}

func GetNroLider() int64{
	return nroLider
}

func GetGroup1() []int64{
	return group1
}

func GetGroup2() []int64{
	return group2
}


// funciones: crea la conexión
func Grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc1)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	lj.RegisterLiderJugadorServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v", port_grpc1)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}
