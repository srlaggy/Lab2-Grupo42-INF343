package main

// LIDER

import (
	// sp "lab/lider/src/sendPlaysNL"
	// rm "lab/lider/src/requestMountPL"
	// sd "lab/lider/src/sendDeadPL"
	// pr "lab/lider/src/playerRecordNL"
	ut "lab/lider/utils"
	rg "lab/lider/src/sendGameLJ"
	sg "lab/lider/src/startGameL"
	e1 "lab/lider/src/EtapaUnoLJ"
	e2 "lab/lider/src/EtapaDosLJ"
	e3 "lab/lider/src/EtapaTresLJ"
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
	time.Sleep(3*time.Second)

	// vivos fin etapa 1
	ut.PrintVivos(sg.GetJugadores(), sg.GetVivos(), sg.GetMaxJug(), 1)

	e2.GroupAndNumberLider()
	fmt.Println("\nEtapa 2\n")
	
	go e2.Grpc_func()
	e2.LogicaEtapaDosAndLoop()
	time.Sleep(3*time.Second)

	// vivos fin etapa 2
	ut.PrintVivos(sg.GetJugadores(), sg.GetVivos(), sg.GetMaxJug(), 2)
	time.Sleep(5*time.Second)
	fmt.Println("\nEtapa 3\n")

	e3.SetterEtapa()
	go e3.Grpc_func()
	e3.LoopAux()
}