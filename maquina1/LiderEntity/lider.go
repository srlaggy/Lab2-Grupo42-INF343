package main

// LIDER

import (
	// sp "lab/lider/src/sendPlaysNL"
	// rm "lab/lider/src/requestMountPL"
	// sd "lab/lider/src/sendDeadPL"
	// pr "lab/lider/src/playerRecordNL"
	rg "lab/lider/src/sendGameLJ"
	sg "lab/lider/src/startGameL"
	e1 "lab/lider/src/EtapaUnoLJ"
	e2 "lab/lider/src/EtapaDosLJ"
	"time"
	"fmt"
)

// --------------- FUNCION MAIN --------------- //

func main(){
	// ----- FUNCIÓN: recibir entradas de jugadores al juego ----- //
	go rg.Grpc_func()

	// Crea arreglo de jugadores para el juego e inicia loop hasta obtener todos los jugadores
	sg.StartGame()

	// ----- FUNCIÓN: pedir monto acumulado al pozo ----- //
	// go rm.RequestMount()

	// ----- FUNCIÓN: enviar jugadas al NameNode ----- //
	// go sp.SendPlaysLider("Jugador_2 Ronda_2 jugada_1")

	// ----- FUNCIÓN: solicitar registro de jugadores al NameNode ----- //
	// go pr.PlayerRecordLider("Jugador_1")

	// ----- FUNCIÓN: enviar los jugadores eliminados al pozo ----- //
	// sd.SendDead_amqp()

	time.Sleep(5*time.Second)
	fmt.Println("\nComienza el juego")
	fmt.Println("Etapa 1\n")

	// funcion que recibe jugadas de jugadores en la primera etapa
	go e1.Grpc_func()
	e1.LoopEtapaUno()
	// time.Sleep(5*time.Second)

	// fmt.Println(sg.GetVivos())

	e2.GroupAndNumberLider()
	fmt.Println("\nEtapa 2\n")
	// fmt.Println("El numero del lider es", e2.GetNroLider())
	// fmt.Println("El grupo 1 es", e2.GetGroup1())
	// fmt.Println("El grupo 2 es", e2.GetGroup2())
	
	go e2.Grpc_func()
	e2.LoopEtapaDos()
	time.Sleep(5*time.Second)

	fmt.Println(sg.GetVivos())
	fmt.Println(e2.GetJugadoresE2())
	fmt.Println(e2.GetVivos())

	time.Sleep(5*time.Second)
}