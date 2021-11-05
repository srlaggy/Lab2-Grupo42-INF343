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

func PrintVivos(jugadores []int, vivos []bool, largo int, etapa int) string{
	var strings string = "Los jugadores vivos al final de la etapa " + strconv.Itoa(etapa) + " son:"
	for i:=0; i<largo; i++{
		if vivos[i]{
			strings += " " + strconv.Itoa(jugadores[i])
			if i+1!=largo{
				strings += ","
			}
		}
	}
	return strings
}