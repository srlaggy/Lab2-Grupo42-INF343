package main

// namenode

import (
	// "fmt"
	aj "lab/namenode/src/administracionJugadas"
)

// --------------- FUNCION PRINCIPAL --------------- //
func main() {

	aj.IniciarRegistroJugadas()

	// servidor grpc que recibe jugada
	/* go sp.Grpc_func()
    pr.Grpc_func() */

	// var num string
	// var ingreso string
	// var delete string

	aj.EntregarJugada("jugador_3 ronda_4 5")
	aj.EntregarJugada("jugador_3 ronda_6 4")
	aj.EntregarJugada("jugador_3 ronda_5 5")
	aj.EntregarJugada("jugador_3 ronda_1 4")
	aj.EntregarJugada("jugador_3 ronda_8 5")
	aj.DevolverJugadasRondas("jugador_3")
	// for{
	// 	fmt.Println(" 1 Enviar jugada \n 2 Pedir jugadas \n 3 Salir \n" )
	// 	fmt.Println("Ingrese un numero:" )
	// 	fmt.Scanln(&num)
		
	// 	if num == "1"{
	// 		fmt.Println("Ingrese jugada:" )
	// 		fmt.Scanln(&ingreso) // Aqui falla FIXME
	// 		fmt.Println(ingreso)
	// 		fmt.Scanln(&delete)
	// 		aj.EntregarJugada(ingreso) // Aqui falla FIXME
	// 	} else if num == "2"{
	// 		fmt.Println("Ingrese jugador:" )
	// 		fmt.Scanln(&ingreso)
	// 		aj.DevolverJugadasRondas(ingreso) // Aqui falla FIXME
	// 	} else if num == "3"{
	// 		break
	// 	}
	// }
}