package administracionJugadas

// Datnode

import (
	// "context"
	// "log"
	// "net"
	"time"
	"fmt"
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
	var fileName string =  "utils/jugadas.txt"
	file, err := os.Create(fileName)
	ut.FailOnError(err, "Failed to create file")
	defer file.Close()
}



// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// FUNCIONES SendPlaysDataNode
// ----------------------------------->

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: elegirDataNode
// -> Dado una combinación de ronda jugador los asigna a un
// DataNode concreto y registra dicho DataNode.
// Recibe: jugador string, ronda string
// Retorna: La dirección ip del dataNode que registrara esa
// combinación de ronda con jugador
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func elegirDataNode(jugador string, ronda string) string {
	Direcciones := [3]string{"Datanode_1", "Datanode_2", "Datanode_3"}
	//TO-DO: Reemplazar por las direcciones ip de los pcs de la u,
	// o las direcciones de prueba segun sea el caso
	var ip string
	rand.Seed(time.Now().UnixNano())
	var seleccionada int= rand.Intn(3)
	ip = Direcciones[seleccionada]
	var lineContent string =  jugador + " " + ronda + " " + ip + "\n"
	var fileName string = "utils/jugadas.txt"
	file, error1 := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	ut.FailOnError(error1, "Failed to open file")
	_, error2 := file.WriteString(lineContent)
	ut.FailOnError(error2, "Failed to write file")
	defer file.Close()
	// ip = "localhost" //DELETE
	return ip
}

/* func iniciarDataNode(jugador string, ronda string){
	var ip string
	ip = elegirDataNode(jugador, ronda)
	//TO-DO: Revisar si se hace así
} */

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: entregarJugada
// -> Con la ayuda de encontrarDataNode obtiene la direccion
// ip para el jugador en la ronda dada y le envia la jugada al
// dataNode correspondiente para su registro.
// Recibe: jugador string, ronda string y jugada como string
// Retorna: jugador string, ronda string y jugada como string
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func EntregarJugada(dato string) {
	
	var ip string = ""
	
	s := strings.Split(dato, " ")
	
	ip = encontrarDataNode(s[0], s[1], false)

	// ip = "localhost" //DELETE
	sps.SendPlaysDataNode(dato, ip)
}



// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// FUNCIONES PlayerRecordNameNode
// ----------------------------------->

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: encontrarDataNode
// -> Localiza la ubicación de una combinación de jugador y
// ronda y retorna la direccion ip del dataNode correspondiente,
// en caso de no existir dicha combinación llama a "elegirDataNode"
// Recibe: jugador string, ronda string y jugada como string
// Retorna: jugador string, ronda string y jugada como string
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
func encontrarDataNode(jugador string, ronda string, flag bool) string{
	// fmt.Println("He entrado a encontrarDataNode", jugador, ronda)
	var ip string= "No hay jugadas"
	fileName := "utils/jugadas.txt"
	file, error1 := os.Open(fileName)
	ut.FailOnError(error1, "Failed to open file")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ubicacion string
	// fmt.Println(scanner)
	for scanner.Scan(){
		ubicacion = scanner.Text()
		s := strings.Split(ubicacion, " ")
		fmt.Println("Jugador:", s[0], "Ronda:", s[1], "Datanode:", s[2])
		if(jugador == s[0]){
			if(ronda == s[1]){
				fmt.Println("Entre!")
				ip = s[2]
				flag = true
			}
		}
    }
	
	if(!flag){
		ip = elegirDataNode(jugador, ronda)
	}
	// ip = "localhost" //DELETE
	return ip
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: SolicitarJugadas
// -> Dada un jugador y una ubicación solicita las jugadas de
// dicho jugador al dataNode.
// Recibe: jugador string, ronda string, ip String
// Retorna: Nada
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


func SolicitarJugadas(jugador string, ronda string) string{
	fmt.Println("He entrado a SolicitarJugadas", "Jugador:", jugador, "Ronda:", ronda)
	ip := encontrarDataNode(jugador, ronda, true)
	fmt.Println("IP de la ronda del jugador:", ip)
	if (ip == "No hay jugadas"){
		fmt.Println("No hay jugada", jugador, ronda)
		return ip
	}
	var jugadas string = ""
	// ip = "localhost" //DELETE
	// fmt.Println("Estamos casi")
	// if (ip == "Datanode_1"){
	jugadas = pr.PlayerRecordNameNode(jugador, ronda, ip)
	// }
	
	
	//TO-DO: hacer un request a dataNode, recibir la request y usarlo
	// en lugar del texto de reemplazo
	jugadas = "veamos"
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
	fmt.Println("\nDEVOLVER JUGADAS RONDAS\n")
	var jugadas string = ""
	var i int = 0
	
	for i < 3{
		i++
		num := strconv.Itoa(i)
		valores := SolicitarJugadas(jugador, "ronda_" + num)
		ronda := "Juego " + num + ": " 
		jugadas = jugadas + ronda + valores + "\n"
	}
	return jugadas
}