package main

// POZO
import (
	sm "lab/pozo/src/sendMountPL"
	rd "lab/pozo/src/receiveDeadPL"
)

// --------------- FUNCION MAIN --------------- //
func main() {
    // reseteamos pozo
    rd.ResetPozo()
    // ----- GRPC ----- //
    // servidor grpc que envia el monto del pozo
    go sm.Grpc_func()

    // ----- RABBITMQ ----- //
    // servidor amqp que recibe un muerto
    rd.Rabbit_func()
}