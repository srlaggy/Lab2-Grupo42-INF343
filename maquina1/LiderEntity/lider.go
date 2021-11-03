package main

// LIDER

import (
	// sp "lab/lider/src/sendPlaysNL"
	// rm "lab/lider/src/requestMountPL"
	// sd "lab/lider/src/sendDeadPL"
	// pr "lab/lider/src/playerRecordNL"
	rg "lab/lider/src/sendGameLJ"
	sg "lab/lider/src/startGameL"
	// e1 "lab/lider/src/EtapaUnoLJ"
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
	
	fmt.Printf("\nComienza el juego\n")

	// funcion que recibe jugadas de jugadores en la primera etapa
	// go e1.Grpc_func()

	// e1.LoopEtapaUno()

	// e2.Group()
	e2.Grpc_func()

	
	// fmt.Println(sg.GetVivos())

	time.Sleep(5*time.Second)
}