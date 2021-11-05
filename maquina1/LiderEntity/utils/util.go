package utils

import (
	"log"
	"strconv"
)

func CreateDir(protocolo string, address string, port string) string{
	if protocolo == ""{
		return address + ":" + port
	} else if address == ""{
		return protocolo + ":" + port
	} else {
		return protocolo + "://" + address + ":" + port
	}
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// largo es len de listas, osea 16
func PrintVivos(jugadores []int, vivos []bool, largo int, etapa int) string{
	var strings string = "Los jugadores vivos al final de la etapa " + strconv.Itoa(etapa) + " son:"
	for i:=0; i<largo; i++{
		if vivos[i]{
			strings += " " + strconv.Itoa(jugadores[i])
		}
	}
	return strings
}

// largo es el numero de vivos no total de jugadores
func PrintWinners(jugadores []int, vivos []bool, largo int) string{
	var strings string
	if largo==0{
		strings = "Por esta vez, no hay ganadores del juego del Calamar"
	} else if largo==1 {
		strings += "El ganador del juego del Calamar es el jugador "
		for i:=0; i<len(jugadores); i++{
			if vivos[i]{
				strings += strconv.Itoa(jugadores[i])
				break
			}
		}
	} else {
		strings += "Los ganadores del juego del Calamar son los jugadores"
		for i:=0; i<len(jugadores); i++{
			if vivos[i]{
				strings += " " + strconv.Itoa(jugadores[i])
			}
		}
	}
	return strings
}