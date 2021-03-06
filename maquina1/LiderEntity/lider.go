package main

// LIDER

import (
	// sp "lab/lider/src/sendPlaysNL"
	rm "lab/lider/src/requestMountPL"
	// sd "lab/lider/src/sendDeadPL"
	lj "lab/lider/proto/LJ"
	pr "lab/lider/src/playerRecordNL"
	ut "lab/lider/utils"
	rg "lab/lider/src/sendGameLJ"
	sg "lab/lider/src/startGameL"
	e1 "lab/lider/src/EtapaUnoLJ"
	e2 "lab/lider/src/EtapaDosLJ"
	e3 "lab/lider/src/EtapaTresLJ"
	"google.golang.org/grpc"
	"time"
	"context"
	"fmt"
	"sync"
	"strconv"
	"net"
)

const (
	protocolo_grpc = "tcp"
	port_grpc = "41000"
)

var monto int64 = 0

type server struct {
	lj.UnimplementedLiderJugadorServiceServer
}

func (s *server) MontoJug(ctx context.Context, in *lj.MontoJugReq) (*lj.MontoJugResp, error) {
	if in.Trigger==55{
		monto = rm.RequestMount()
	}
	return &lj.MontoJugResp{MontoJugador: monto}, nil
}

// --------------- FUNCION PRINCIPAL --------------- //
func Grpc_func_pozo() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	lj.RegisterLiderJugadorServiceServer(s, &server{})
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}

// funcion para crear interfaz
func interfaz(wg *sync.WaitGroup){
	defer wg.Done()
	var eleccion int64
	fmt.Println("\nEntretiempo")
	fmt.Println("Tomate un descanso")
	fmt.Println("Selecciona una de las opciones:")
	fmt.Println("1) Presione 1 para consultar las jugadas de un jugador")
	fmt.Println("2) Presione 2 para continuar a la siguiente etapa\n")
	fmt.Scanln(&eleccion)
	time.Sleep(2*time.Second)
	for (eleccion!=1 && eleccion!=2){
		fmt.Println("\nVuelva a ingresar una opcion")
		fmt.Scanln(&eleccion)
	}
	if eleccion==1{
		var eleccion2 int64
		fmt.Println("\nIngrese Numero de jugador a consultar")
		fmt.Println("Recuerde que existen 16 jugadores\n")
		fmt.Scanln(&eleccion2)
		time.Sleep(2*time.Second)
		for (eleccion2<1 || eleccion2>16){
			fmt.Println("\nVuelva a ingresar el numero del jugador")
			fmt.Println("Recuerde que existen 16 jugadores\n")
			fmt.Scanln(&eleccion2)
		}
		strings := pr.PlayerRecordLider("Jugador_" + strconv.Itoa(int(eleccion2)))
		fmt.Println("Imprimiendo Jugadas")
		fmt.Println(strings)
		time.Sleep(3*time.Second)
		fmt.Println("Ahora avanzamos a la siguiente etapa")
	} else if eleccion==2{
		fmt.Println("Pongase comodo, que avanzamos a la siguiente etapa")
	}
}

// --------------- FUNCION MAIN --------------- //

func main(){
	// ----- FUNCI??N: recibir entradas de jugadores al juego ----- //
	go rg.Grpc_func()

	// server de pedir monto al pozo
	go Grpc_func_pozo()

	// Crea arreglo de jugadores para el juego e inicia loop hasta obtener todos los jugadores
	sg.StartGame()

	// ----- FUNCI??N: enviar los jugadores eliminados al pozo ----- //
	// sd.SendDead_amqp()

	time.Sleep(5*time.Second)
	fmt.Println("\nComienza el juego")
	fmt.Println("Etapa 1\n")

	// funcion que recibe jugadas de jugadores en la primera etapa
	go e1.Grpc_func()
	e1.LoopEtapaUno()
	time.Sleep(3*time.Second)

	// vivos fin etapa 1
	fmt.Println(ut.PrintVivos(sg.GetJugadores(), sg.GetVivos(), sg.GetMaxJug(), 1))

	// interfaz 1-2
	var wg sync.WaitGroup
	wg.Add(1)
	go interfaz(&wg)
	wg.Wait()

	e2.GroupAndNumberLider()
	fmt.Println("\nEtapa 2\n")
	
	go e2.Grpc_func()
	e2.LogicaEtapaDosAndLoop()
	time.Sleep(3*time.Second)

	// vivos fin etapa 2
	fmt.Println(ut.PrintVivos(sg.GetJugadores(), sg.GetVivos(), sg.GetMaxJug(), 2))

	// interfaz 2-3
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go interfaz(&wg2)
	wg2.Wait()

	time.Sleep(5*time.Second)
	fmt.Println("\nEtapa 3\n")

	e3.SetterEtapa()
	go e3.Grpc_func()
	e3.LoopAux()

	fmt.Println(ut.PrintWinners(sg.GetJugadores(), sg.GetVivos(), e3.GetVivosTotales()))
	time.Sleep(3*time.Second)
}