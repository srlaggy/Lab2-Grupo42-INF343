package main

import "os"

// Datnode

/* import(
    "context"
    "log"
    "net"
    "time"
    "fmt"
	"math/rand"
    "google.golang.org/grpc"
	"bufio"
	"os"
	"strconv"
) */

func main() {
    
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: iniciarRegistroJugadas
// -> Crea/limpia el registro de ubicación de las jugadas
// Recibe: Nada
// Retorna: Nada
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
  
func iniciarRegistroJugadas(){
	//Se crea el archivo
	var fileName string =  "jugadas.txt"
	file, err := os.Create(fileName)
	//failOnError(err, "Failed to create file")
	//FIXME: undeclared name: failOnError
	defer file.Close()
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
	Direcciones := [3]int{"1", "2", "3"}
	//TO-DO: Reemplazar por las direcciones ip de los pcs de la u,
	// o las direcciones de prueba segun sea el caso
	
	var seleccionada int= rand.Intn(3)
	ip = Direcciones[seleccionada]
	var lineContent string =  jugador + " " + ronda + " " + ip + "\n"
	var fileName string = "jugadas.txt"
	file, error1 := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	_, error2 := file.WriteString(lineContent)
	return ip
}

func iniciarDataNode(jugador string, ronda string){
	ip = elegirDataNode(jugador string, ronda string)
	//TO-DO: Revisar si se hace así
}


// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: entregarJugada
// -> Localiza la ubicación de una combinación de jugador y
// ronda y retorna la direccion ip del dataNode correspondiente,
// en caso de no existir dicha combinación llama a "elegirDataNode"
// Recibe: jugador string, ronda string y jugada como string
// Retorna: jugador string, ronda string y jugada como string
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func encontrarDataNode(jugador string, ronda string) string ip{
	var ip string
	file, error1 := os.OpenFile("jugadas.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	scanner := bufio.NewScanner(f)
	scanner := bufio.NewScanner(file)
	flag = false
    for scanner.Scan(){
		ubicacion = scanner.Text()
		s := strings.Fields(ubicacion)
		if(jugador == s[0]){
			if(ronda == s[1]){
				ip = s[2]
				flag = true
			}
		}

    }
	if(!flag){
		ip = elegirDataNode(jugador, ronda)
	}
	return ip
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: entregarJugada
// -> Con la ayuda de encontrarDataNode obtiene la direccion
// ip para el jugador en la ronda dada y le envia la jugada al
// dataNode correspondiente para su registro.
// Recibe: jugador string, ronda string y jugada como string
// Retorna: jugador string, ronda string y jugada como string
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func entregarJugada(jugador string, ronda string, jugada string)(jugador string, ronda string, jugada string){
	var lineContent string =  jugador + " " + ronda
	var fileName string = "jugadas.txt"
	file, error1 := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	ip = encontrarDataNode(jugador, ronda)

	return jugador, ronda, jugada

}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: solicitarJugadasRonda
// -> Dada un jugador y una ubicación solicita las jugadas de
// dicho jugador al dataNode.
// Recibe: jugador string, ronda string, ip String
// Retorna: Nada
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


func solicitarJugadasRondas(jugador string, ronda string) int []{
		ip = encontrarDataNode(jugador, i)
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: devolverJugadasRondas
// -> Localiza la ubicación de todos los registros de un jugador
// dado para así lograr comunicar todas las jugadas de un jugador
// al líder.
// Recibe: jugador string
// Retorna: todas las jugadas de un jugador
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func devolverJugadasRondas(jugador string) string{
	
	for i < 3{
		i++
		solicitarJugadasRondas(jugador, i)
	}
}

