package main

// namenode

import (
	"fmt"
	aj "lab/namenode/src/administracionJugadas"
)

// --------------- FUNCION PRINCIPAL --------------- //
func main() {

	aj.IniciarRegistroJugadas()

	// servidor grpc que recibe jugada
	/* go sp.Grpc_func()
    pr.Grpc_func() */

	var num string
	var ingreso string
	var valorTemp string

	for{
		fmt.Println("1 Enviar jugada \n 2 Pedir jugadas \n 3 Salir \n" )
		fmt.Println("Ingrese un numero: \n" )
		fmt.Scanln(&num)
		
		if num == "1"{
			fmt.Println("Ingrese jugada: \n" )
			fmt.Scanln(&ingreso) //Aqui se cierra ¬¬ FIXME
			fmt.Println("Recibida\n" )
			fmt.Scanln(&valorTemp)
			aj.EntregarJugada(ingreso)
		} else if num == "2"{
			fmt.Println("Ingrese jugador: \n" )
			fmt.Scanln(&ingreso)
			aj.DevolverJugadasRondas(ingreso)
		} else if num == "3"{
			break
		}
	}
	


}