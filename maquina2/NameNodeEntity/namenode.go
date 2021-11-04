package main

// namenode

import(
	sp "lab/namenode/src/sendPlaysNL"
    pr "lab/namenode/src/playerRecordNL"
	spd "lab/namenode/src/sendPlaysDN"
    prd "lab/namenode/src/playerRecordDN"
	aj "lab/namenode/src/AdministracionJugadas"
)

// --------------- FUNCION PRINCIPAL --------------- //
func main() {

	aj.iniciarRegistroJugadas()

	// servidor grpc que recibe jugada
	go sp.Grpc_func()
    pr.Grpc_func()
}