package main

// POZO
import (
	sm "lab/pozo/src/sendMountPL"
	rd "lab/pozo/src/receiveDeadLP"
)

// --------------- FUNCION MAIN --------------- //
func main() {
    // ----- GRPC ----- //
    go sm.Grpc_func()

    // ----- RABBITMQ ----- //
    rd.Rabbit_func()
}