package startGameL

import (
	// "fmt"
)

const (
	maxJug = 3 // cambiar a 16
)

// VARIABLES GLOBALES
var jugadores []int
var change bool = false

func StartGame(){
	jugadores = make([]int, 0, maxJug)

	for len(jugadores)<=maxJug{
		if change{
			// fmt.Printf("\nNúmero de jugadores: %d\n", GetCantidadJugadores())
			ChangeJugadores()
			if len(jugadores)==maxJug{
				break
			}
		}
	}
}

func AddPlayerGame() int64{
	if len(jugadores)<cap(jugadores){
		jugadores = append(jugadores, len(jugadores)+1)
		ChangeJugadores()
		return int64(len(jugadores))
	} else {
		return 0
	}
}

func GetCantidadJugadores() int{
	return len(jugadores)
}

func ChangeJugadores() {
	if change==false{
		change = true
	} else {
		change = false
	}
}