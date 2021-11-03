package EtapaUnoLJ

import (
	"context"
	"log"
	"net"
	"fmt"
	"math/rand"
	"time"
	"google.golang.org/grpc"
	lj "lab/lider/proto/LJ"
	ut "lab/lider/utils"
	sg "lab/lider/src/startGameL"
)

const (
	address = "localhost"
	protocolo_grpc = "tcp"
	port_grpc1 = "50001"
	min = 6
	max = 10
)

// variable global
var nroLider []int

// ----- FUNCIÓN: recibir jugadas del jugador ----- // --> Lider actua como servidor
type server struct {
	lj.UnimplementedLiderJugadorServiceServer
}

// funcion: primer juego
func (s *server) EtapaUno(ctx context.Context, in *lj.NumPasosReq) (*lj.NumPasosResp, error) {
	// 2 muerto por rondas - 1 vivo - 0 muerto
	var msg int64 = 1
	if in.PlayMsg >= int64(nroLider[in.Ronda]){
		sg.SetVivos(int(in.NroJugador)-1, false)
		msg = 0
		fmt.Printf("El jugador %d ha muerto\n", in.NroJugador)
	}else if in.Ronda==3 && in.Total<21{
		sg.SetVivos(int(in.NroJugador)-1, false)
		msg = 2
		fmt.Printf("El jugador %d ha muerto\n", in.NroJugador)
	}
	return &lj.NumPasosResp{StateMsg: msg}, nil
}

// funciones: crea la conexión
func Grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc1)
	ut.FailOnError(err, "Failed to listen")

	nroLider = CreateRandomNumbers()
	fmt.Printf("\nel numero del lider en ronda 1 es %d\n", nroLider[0])
	fmt.Printf("\nel numero del lider en ronda 2 es %d\n", nroLider[1])
	fmt.Printf("\nel numero del lider en ronda 3 es %d\n", nroLider[2])
	fmt.Printf("\nel numero del lider en ronda 4 es %d\n", nroLider[3])

	s := grpc.NewServer()
	lj.RegisterLiderJugadorServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v", port_grpc1)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}

// FUNCIONES AUXILIARES

// determina numero random del lider
func RandomNumber() int64{
	rand.Seed(time.Now().UnixNano())
	aux := rand.Intn(max - min) + min
	return int64(aux)
}

// crear numeros randoms para las 4 rondas
func CreateRandomNumbers() []int{
	aux := make([]int, 4)
	for i:=0; i<4; i++{
		aux[i] = int(RandomNumber())
	}
	return aux
}

// funcion que revisa si todos murieron
func RevVivos() bool{
	aux := sg.GetVivos()
	// retorna si todos estan muertos
	muertos := true
	for i:=0; i<len(aux); i++{
		// si pillo uno que no este muerto, en verdad no estan muertos todos
		if aux[i]{
			muertos = false
		}
	}
	// true -> todos muertos
	// false -> NO todos muertos
	return muertos
}

// bucle de etapa 1
func LoopEtapaUno(){
	for{
		if RevVivos(){
			break
		}
	}
}