package main

// namenode

import(
	sp "lab/namenode/src/sendPlaysNL"
    pr "lab/namenode/src/playerRecordNL"
)

// --------------- FUNCION PRINCIPAL --------------- //
func main() {
	// servidor grpc que recibe jugada
	go sp.Grpc_func()
    pr.Grpc_func()
}