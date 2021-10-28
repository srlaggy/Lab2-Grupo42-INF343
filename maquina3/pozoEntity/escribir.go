package main

import (
    "log"
    "os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func writePozo(dead string){
	fileName := "pozo.txt"
	// Abrimos archivo en modo append
    f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	failOnError(err, "Failed to open file")
	defer f.Close()
	
	// string formateado
	newDead := dead + "\n"
	// Escribimos en archivo
	_, err := f.WriteString(newDead)
	failOnError(err, "Failed to write file")
}

func getMountPozo() int{
	f, err := os.Open("pozo.txt")
	failOnError(err, "Failed to open file")
    defer f.Close()

    scanner := bufio.NewScanner(f)
	var auxLine string
    for scanner.Scan(){
		auxLine = scanner.Text()
    }
	failOnError(scanner.Err(), "Failed to read file")

	mount := strings.Split(auxLine, " ")[2]
	mountInt,err := strconv.Atoi(mount)
	failOnError(err, "Failed to transform to int")
	return mountInt
}

func resetPozo(){
	f, err := os.Create("pozo.txt")
	failOnError(err, "Failed to create file")
    defer f.Close()
	
	_, err2 := f.WriteString("Jugador_0 Ronda_0 0\n")
	failOnError(err2, "Failed to write file")
}

func main() {
    // aux := "Jugador_2 Ronda_2 20123456"
	// writePozo(aux)
	// resetPozo()
	monto := getMountPozo()
	fmt.Println(monto + 20)
}