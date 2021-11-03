package main

import (
	rg "lab/jugador/src/requestGameLJ"
	e1 "lab/jugador/src/EtapaUnoLJ"
)

// --------------- FUNCION MAIN --------------- //

func main(){

	// ----- FUNCIÓN: pedir ingresar al juego ----- //
	rg.RequestGame("Entrar")

	// Inicio Juego 1
	e1.StartGameUno()
}