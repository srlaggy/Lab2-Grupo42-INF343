package main

// namenode

import (
	// "fmt"
	aj "lab/namenode/src/administracionJugadas"
	sp "lab/namenode/src/sendPlaysNL"
	pr "lab/namenode/src/playerRecordDN"
)

// --------------- FUNCION PRINCIPAL --------------- //
func main() {

	aj.IniciarRegistroJugadas()

	// servidor grpc que recibe jugada
	sp.Grpc_func()

	// Devolver el registro de jugadas del jugador con el 
    // pr.Grpc_func()

	// var num string
	// var ingreso string
	// var delete string

	// aj.EntregarJugada("jugador_3 ronda_1 5")
	// aj.EntregarJugada("jugador_3 ronda_2 4")
	// aj.EntregarJugada("jugador_4 ronda_3 5")
	// aj.EntregarJugada("jugador_2 ronda_1 4")
	// aj.EntregarJugada("jugador_3 ronda_2 5")
	// aj.DevolverJugadasRondas("jugador_3")
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