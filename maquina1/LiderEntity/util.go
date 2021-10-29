package main

import (
	"log"
)

func createDir(protocolo string, address string, port string) string{
	if protocolo == ""{
		return address + ":" + port
	} else if address == ""{
		return protocolo + ":" + port
	} else {
		return protocolo + "://" + address + ":" + port
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}