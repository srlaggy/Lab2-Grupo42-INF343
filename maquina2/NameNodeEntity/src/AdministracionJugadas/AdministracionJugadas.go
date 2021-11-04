package namenode

// Datnode

import (
	/* "context"
	"log"
	"net"
	"time"
	"fmt"
	"google.golang.org/grpc" */
	"bufio"
	pr "lab/namenode/src/playerRecordDN"
	sps "lab/namenode/src/sendPlaysDN"
	ut "lab/namenode/utils"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
    
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: iniciarRegistroJugadas
// -> Crea/limpia el registro de ubicación de las jugadas
// Recibe: Nada
// Retorna: Nada
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
  
func IniciarRegistroJugadas(){
	//Se crea el archivo
	var fileName string =  "jugadas.txt"
	file, err := os.Create(fileName)
	ut.FailOnError(err, "Failed to create file")
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
	Direcciones := [3]string{"1", "2", "3"}
	//TO-DO: Reemplazar por las direcciones ip de los pcs de la u,
	// o las direcciones de prueba segun sea el caso
	var ip string
	var seleccionada int= rand.Intn(3)
	ip = Direcciones[seleccionada]
	var lineContent string =  jugador + " " + ronda + " " + ip + "\n"
	var fileName string = "jugadas.txt"
	file, error1 := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	ut.FailOnError(error1, "Failed to open file")
	_, error2 := file.WriteString(lineContent)
	ut.FailOnError(error2, "Failed to write file")
	defer file.Close()

	ip = "localhost" //DELETE
	return ip
}

func iniciarDataNode(jugador string, ronda string){
	var ip string
	ip = elegirDataNode(jugador, ronda)
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
func encontrarDataNode(jugador string, ronda string, flag bool) string{
	var ip string= "No hay jugadas"
	file, error1 := os.OpenFile("jugadas.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	ut.FailOnError(error1, "Failed to open file")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var ubicacion string
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
	ip = "localhost" //DELETE
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

func entregarJugada(dato string) {
	
	var ip string

	s := strings.Fields(dato)
	ip = encontrarDataNode(s[0], s[1], false)

	ip = "localhost" //DELETE
	sps.SendPlaysDataNode(dato, ip)
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: SolicitarJugadas
// -> Dada un jugador y una ubicación solicita las jugadas de
// dicho jugador al dataNode.
// Recibe: jugador string, ronda string, ip String
// Retorna: Nada
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


func SolicitarJugadas(jugador string, ronda string) string{
		
	ip := encontrarDataNode(jugador, ronda, true)
	if(ip == "No hay jugadas"){
		return ip
	}

	var jugadas string = ""

	ip = "localhost" //DELETE
	
	jugadas = pr.PlayerRecordNameNode(jugador, ronda, ip)
	

	//TO-DO: hacer un request a dataNode, recibir la request y usarlo
	// en lugar del texto de reemplazo

	return jugadas
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: devolverJugadasRondas
// -> Localiza la ubicación de todos los registros de un jugador
// dado para así lograr comunicar todas las jugadas de un jugador
// al líder.
// Recibe: jugador string
// Retorna: todas las jugadas de un jugador
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func DevolverJugadasRondas(jugador string) string{
	var jugadas string = ""
	var i int = 0
	
	for i < 3{
		i++
		num := strconv.Itoa(i)
		valores := SolicitarJugadas(jugador, num)
		ronda := "Juego " + num + ": " 
		jugadas = jugadas + ronda + valores + "\n"
	}
	return jugadas
}