package main

// namenode

import (
	"fmt"
	aj "lab/namenode/src/administracionJugadas"
	pr "lab/namenode/src/playerRecordNL"
	sp "lab/namenode/src/sendPlaysNL"
)

// --------------- FUNCION PRINCIPAL --------------- //
func main() {

	aj.iniciarRegistroJugadas()

	// servidor grpc que recibe jugada
	go sp.Grpc_func()
    pr.Grpc_func()

	var num string
	var ingreso string

	for{
		fmt.Println("1 Enviar jugada \n 2 Pedir jugadas \n 3 Salir \n" )
		fmt.Println("Ingrese un numero: \n" )
		fmt.Scanln(&num)
		
		if num == "1"{
			fmt.Scanln(&ingreso)
			aj.EntregarJugada(ingreso)
		} else if num == "2"{
			fmt.Scanln(&ingreso)
			aj.DevolverJugadasRondas(ingreso)
		} else if num == "3"{
			break
		}
	}
	


}