package main

// Datnode

func main() {
    
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: iniciarRegistroJugadas
// -> Crea/limpia el registro de ubicación de las jugadas
// Recibe: Nada
// Retorna: Nada
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
  
func iniciarRegistroJugadas(){

}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: elegirDataNode
// -> Dado una combinación de ronda jugador los asigna a un
// DataNode concreto y registra dicho DataNode.
// Recibe: jugador string, ronda string
// Retorna: La dirección ip del dataNode que registrara esa
// combinación de ronda con jugador
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func elegirDataNode(jugador string, ronda string) string {
	ip = "0"
	return ip
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: entregarJugada
// -> Localiza la ubicación de una combinación de jugador y
// ronda y envia la jugada a ese dataNode, en caso de no existir
// dicha combinación llama a "elegirDataNode"
// Recibe: jugador string, ronda string y jugada como string
// Retorna: Nada
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func entregarJugada(jugador string, ronda string, jugada string){

}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: solicitarJugadasRonda
// -> Dada un jugador y una ubicación solicita las jugadas de
// dicho jugador al dataNode.
// Recibe: jugador string, ronda string, ip String
// Retorna: Nada
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


func solicitarJugadasRonda(jugador string, ronda string, ip string) {}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: devolverJugadasRondas
// -> Localiza la ubicación de todos los registros de un jugador
// dado para así lograr comunicar todas las jugadas de un jugador
// al líder.
// Recibe: jugador string
// Retorna: todas las jugadas de un jugador
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func devolverJugadasRondas(jugador string) []int{}

