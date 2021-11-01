package main

// namenode

import(
	sp "lab/namenode/src/sendPlaysNL"
)

// --------------- FUNCION PRINCIPAL --------------- //
func main() {
	// servidor grpc que recibe jugada
	sp.Grpc_func()
}