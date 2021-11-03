package main

import (
	rg "lab/jugador/src/requestGameLJ"
	// e1 "lab/jugador/src/EtapaUnoLJ"
	e2 "lab/jugador/src/EtapaDosLJ"
)

// --------------- FUNCION MAIN --------------- //

func main(){

	// ----- FUNCIÓN: pedir ingresar al juego ----- //
	rg.RequestGame("Entrar")

	// Inicio Juego 1
	// e1.StartGameUno()
	e2.Etapa2Conn()
	if (e2.GetVivo_muerto()){
		e2.Etapa2()
	}
}