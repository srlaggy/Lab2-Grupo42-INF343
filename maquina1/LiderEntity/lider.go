package main

// LIDER

import (
	sp "lab/lider/src/sendPlaysNL"
	rm "lab/lider/src/requestMountPL"
	sd "lab/lider/src/sendDeadPL"
)

const (
	// address = "localhost"
	// protocolo_grpc = ""
	// port_grpc2 = "60001"
)

// --------------- FUNCION MAIN --------------- //

func main(){

	// ----- FUNCIÓN: pedir monto acumulado al pozo ----- //
	go rm.RequestMount()

	// ----- FUNCIÓN: enviar jugadas al NameNode ----- //
	go sp.SendPlaysLider("Jugador_2 Ronda_2 jugada_1")

	// ----- FUNCIÓN: enviar los jugadores eliminados al pozo ----- //
	sd.SendDead_amqp()
}