package main

import (
	rg "lab/jugador/src/requestGameLJ"
	e1 "lab/jugador/src/EtapaUnoLJ"
	e2 "lab/jugador/src/EtapaDosLJ"
	e3 "lab/jugador/src/EtapaTresLJ"
	ut "lab/jugador/utils"
	lj "lab/jugador/proto/LJ"
	"google.golang.org/grpc"
	"fmt"
	"time"
	"sync"
	"context"
)

const (
	address = "localhost"
	protocolo_grpc = ""
	port_grpc = "41000"
)

func RequestMountJug(num int64) int64{
	// Set up a connection to the server.
	conn1, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn1.Close()

	c := lj.NewLiderJugadorServiceClient(conn1)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.MontoJug(ctx, &lj.MontoJugReq{Trigger: num})
	ut.FailOnError(err, "Failed to send a mount")
	return int64(r.GetMontoJugador())
}

// funcion para crear interfaz
func interfaz(wg *sync.WaitGroup){
	defer wg.Done()
	var eleccion int64
	fmt.Println("\nEntretiempo")
	fmt.Println("Superaste la etapa y mereces un descanso")
	fmt.Println("Selecciona una de las opciones:")
	fmt.Println("1) Presione 1 para consultar el monto acumulado del pozo")
	fmt.Println("2) Presione 2 para continuar a la siguiente etapa\n")
	fmt.Scanln(&eleccion)
	time.Sleep(2*time.Second)
	for (eleccion!=1 && eleccion!=2){
		fmt.Println("\nVuelva a ingresar una opcion")
		fmt.Scanln(&eleccion)
	}
	if eleccion==1{
		aux := RequestMountJug(55)
		time.Sleep(1*time.Second)
		fmt.Println("\nEl monto acumulado en el pozo es de", aux, "wones")
		time.Sleep(3*time.Second)
		fmt.Println("Ahora avanzamos a la siguiente etapa")
	} else if eleccion==2{
		fmt.Println("Preparate!! La siguiente etapa empieza en breve")
	}
}

// --------------- FUNCION MAIN --------------- //

func main(){

	// ----- FUNCIÓN: pedir ingresar al juego ----- //
	rg.RequestGame("jugador1 - humano")
	rg.RequestGameBot("Jugador2 - Bot")
	rg.RequestGameBot("Jugador3 - Bot")
	rg.RequestGameBot("Jugador4 - Bot")
	rg.RequestGameBot("Jugador5 - Bot")
	rg.RequestGameBot("Jugador6 - Bot")
	rg.RequestGameBot("Jugador7 - Bot")
	rg.RequestGameBot("Jugador8 - Bot")
	rg.RequestGameBot("Jugador9 - Bot")
	rg.RequestGameBot("Jugador10 - Bot")
	rg.RequestGameBot("Jugador11 - Bot")
	rg.RequestGameBot("Jugador12 - Bot")
	rg.RequestGameBot("Jugador13 - Bot")
	rg.RequestGameBot("Jugador14 - Bot")
	rg.RequestGameBot("Jugador15 - Bot")
	rg.RequestGameBot("Jugador16 - Bot")

	time.Sleep(5*time.Second)

	// Inicio Juego 1
	e1.StartGameUnoTrigger()

	// interfaz etapa 1-2
	if rg.GetVivosSlice()[0]{
		var wg sync.WaitGroup
		wg.Add(1)
		go interfaz(&wg)
		wg.Wait()
	}

	// Inicio Juego 2
	time.Sleep(5*time.Second)
	e2.Etapa2ConnTrigger()
	e2.StartEtapa2Trigger()

	// Inicio Juego 3
	time.Sleep(5*time.Second)
	e3.Etapa3ConnTrigger()
	time.Sleep(2*time.Second)
	e3.Etapa3Trigger()
	time.Sleep(10*time.Second)

	fmt.Println(e3.GetJugadoresE3())
	fmt.Println(e3.GetVivosE3())
	fmt.Println(e3.GetCantVivos())

	// fmt.Println(e2.GetVivosNumerosFinal())
}