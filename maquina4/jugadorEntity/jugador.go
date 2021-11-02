package main

import (
	rg "lab/jugador/src/requestGameLJ"
	sp "lab/jugador/src/sendPlaysLJ"
)

// --------------- FUNCION MAIN --------------- //

func main(){

	// ----- FUNCIÓN: pedir ingresar al juego ----- //
	rg.RequestGame("Entrar")

	sp.SendPasos()
}