package main

// namenode

import (
	// "fmt"
	aj "lab/namenode/src/administracionJugadas"
	sp "lab/namenode/src/sendPlaysNL"
	pr "lab/namenode/src/playerRecordNL"
)

// --------------- FUNCION PRINCIPAL --------------- //
func main() {
	aj.IniciarRegistroJugadas()

	// servidor grpc que recibe jugada
	go sp.Grpc_func()

	// Devolver el registro de jugadas del jugador con el 
    pr.Grpc_func()
}