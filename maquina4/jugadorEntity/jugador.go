package main

import (
	rg "lab/jugador/src/requestGameLJ"
	e1 "lab/jugador/src/EtapaUnoLJ"
	e2 "lab/jugador/src/EtapaDosLJ"
	"fmt"
)

// --------------- FUNCION MAIN --------------- //

func main(){

	// ----- FUNCIÓN: pedir ingresar al juego ----- //
	rg.RequestGame("jugador1 - humano")
	rg.RequestGameBot("Jugador2 - Bot")
	rg.RequestGameBot("Jugador3 - Bot")
	rg.RequestGameBot("Jugador4 - Bot")
	rg.RequestGameBot("Jugador5 - Bot")
	rg.RequestGameBot("Jugador6 - Bot")
	rg.RequestGameBot("Jugador7 - Bot")
	rg.RequestGameBot("Jugador8 - Bot")
	rg.RequestGameBot("Jugador9 - Bot")
	rg.RequestGameBot("Jugador10 - Bot")
	rg.RequestGameBot("Jugador11 - Bot")
	rg.RequestGameBot("Jugador12 - Bot")
	rg.RequestGameBot("Jugador13 - Bot")
	rg.RequestGameBot("Jugador14 - Bot")
	rg.RequestGameBot("Jugador15 - Bot")
	rg.RequestGameBot("Jugador16 - Bot")

	// Inicio Juego 1
	e1.StartGameUnoTrigger()

	// Inicio Juego 2
	e2.Etapa2ConnTrigger()
	e2.StartEtapa2Trigger()

	// fmt.Println(rg.GetVivosSlice())
	fmt.Println(e2.GetVivosNumeros())
	fmt.Println(e2.GetVivosBool())
}