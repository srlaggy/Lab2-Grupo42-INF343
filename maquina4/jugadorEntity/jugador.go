package main

import (
	rg "lab/jugador/src/requestGameLJ"
	e1c "lab/jugador/src/EtapaUnoLJ"
)

// --------------- FUNCION MAIN --------------- //

func main(){

	// ----- FUNCIÓN: pedir ingresar al juego ----- //
	rg.RequestGame("Entrar")

	e1c.EtapaUno()
}