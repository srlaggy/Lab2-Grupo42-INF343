package main

// LIDER

import (
	sp "lab/lider/src/sendPlaysNL"
	rm "lab/lider/src/requestMountPL"
	sd "lab/lider/src/sendDeadPL"
	pr "lab/lider/src/playerRecordNL"
	rg "lab/lider/src/sendGameLJ"
	"time"
)

// --------------- FUNCION MAIN --------------- //

func main(){

	// ----- FUNCIÓN: pedir monto acumulado al pozo ----- //
	go rm.RequestMount()

	// ----- FUNCIÓN: enviar jugadas al NameNode ----- //
	go sp.SendPlaysLider("Jugador_2 Ronda_2 jugada_1")

	// ----- FUNCIÓN: solicitar registro de jugadores al NameNode ----- //
	go pr.PlayerRecordLider("Jugador_1")

	// ----- FUNCIÓN: recibir entradas de jugadores al juego ----- //
	go rg.Grpc_func()

	// ----- FUNCIÓN: enviar los jugadores eliminados al pozo ----- //
	sd.SendDead_amqp()

	time.Sleep(5*time.Second)
}