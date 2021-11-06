package EtapaUnoLJ

import (
	"context"
	"strconv"
	"net"
	"fmt"
	"math/rand"
	"time"
	"google.golang.org/grpc"
	lj "lab/lider/proto/LJ"
	ut "lab/lider/utils"
	sg "lab/lider/src/startGameL"
	sp "lab/lider/src/sendPlaysNL"
	sd "lab/lider/src/sendDeadPL"
)

const (
	address = "localhost"
	protocolo_grpc = "tcp"
	port_grpc1 = "50001"
	// min = 6   // descomentar
	// max = 10  // descomentar
	min = 11
	max = 12
)

// variable global
var nroLider []int
var juegoUno []bool

// ----- FUNCIÓN: recibir jugadas del jugador ----- // --> Lider actua como servidor
type server struct {
	lj.UnimplementedLiderJugadorServiceServer
}

// funcion: primer juego
func (s *server) EtapaUno(ctx context.Context, in *lj.NumPasosReq) (*lj.NumPasosResp, error) {
	// 3 gano el juego (sin importar la ronda) - 2 muerto por rondas - 1 vivo - 0 muerto
	var msg int64 = 1
	// Se llama a la funcion sendPlaysNL para que el lider envie su jugada al NameNode
    // ejemplo de jugada: "Jugador_2 Ronda_2 jugada_1"
    sp.SendPlaysLider("Jugador_"+strconv.FormatInt(in.NroJugador,10)+" Ronda_1 "+strconv.FormatInt(in.PlayMsg,10))
	if in.PlayMsg >= int64(nroLider[in.Ronda]){
		sg.SetVivos(int(in.NroJugador)-1, false)
		// ya jugo el juego
		juegoUno[int(in.NroJugador)-1] = true
		msg = 0
		fmt.Printf("El jugador %d ha muerto\n", in.NroJugador)
		// se informa al pozo de la defuncion
		sd.SendDead_amqp(sd.Muertos(int(in.NroJugador), 1))
	}else if in.Ronda==3 && in.Total+in.PlayMsg<21{
		sg.SetVivos(int(in.NroJugador)-1, false)
		// ya jugo el juego
		juegoUno[int(in.NroJugador)-1] = true
		msg = 2
		fmt.Printf("El jugador %d ha muerto\n", in.NroJugador)
		// se informa al pozo de la defuncion
		sd.SendDead_amqp(sd.Muertos(int(in.NroJugador), 1))
	}else if in.Total+in.PlayMsg>=21{
		msg = 3
		// ya jugo el juego
		juegoUno[int(in.NroJugador)-1] = true
		fmt.Printf("El jugador %d supero el juego 1 en la ronda %d\n", in.NroJugador, in.Ronda + 1)
	}
	return &lj.NumPasosResp{StateMsg: msg}, nil
}

// funciones: crea la conexión
func Grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc1)
	ut.FailOnError(err, "Failed to listen")

	nroLider = CreateRandomNumbers()

	s := grpc.NewServer()
	lj.RegisterLiderJugadorServiceServer(s, &server{})
	// log.Printf("Servidor grpc escuchando en el puerto %v", port_grpc1)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}

// FUNCIONES AUXILIARES

// determina numero random
func RandomNumber() int64{
	rand.Seed(time.Now().UnixNano())
	aux := rand.Intn(max - min) + min
	return int64(aux)
}

// crear numeros randoms para las 4 rondas
func CreateRandomNumbers() []int{
	rand.Seed(time.Now().UnixNano())
	aux := make([]int, 4)
	for i:=0; i<4; i++{
		aux[i] = rand.Intn(max - min) + min
	}
	return aux
}

// crear numeros randoms para las 4 rondas
func CreateJugJuegoUno() []bool{
	aux := make([]bool, sg.GetMaxJug())
	for i:=0; i<len(aux); i++{
		aux[i] = false
	}
	return aux
}

// funcion que revisa si todos jugaron
func RevJugaron() bool{
	// retorna si todos jugaron
	jugaron := true
	for i:=0; i<len(juegoUno); i++{
		// si pillo uno que no haya jugado, en verdad no han jugado todos
		if !juegoUno[i]{
			jugaron = false
		}
	}
	// true -> todos jugaron
	// false -> NO todos jugaron
	return jugaron
}

// bucle de etapa 1
func LoopEtapaUno(){
	juegoUno = CreateJugJuegoUno()
	for{
		if RevJugaron(){
			break
		}
	}
}