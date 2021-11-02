package main

// Datnode

import (
	/* "context"
	   "log"
	   "net"
	   "time"
	   "fmt"
	   "google.golang.org/grpc" */
	"bufio"
	"os"
	"strings"
)

func main() {
    
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: crearRegistroJugadas
// -> Genera el registro de un jugador por primera vez en la
// ronda
// Recibe: el jugador y la ronda, ambos como string
// Retorna: Nada
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
  
func crearRegistroJugadas(registro string){

	//TO-DO: Redefinir el input en caso de ser necesario
	//TO-DO: En caso de redefinir el input hacer de ronda y
	//jugador textos
	//NOTE: jugador_n__ronda_m.txt
	
	//Se crea el archivo
	s := strings.Fields(dato)
	var fileName string = s[0] + "__" + s[1] + ".txt"
	file, err := os.Create(fileName)
	//FailOnError(err, "Failed to create file")
	//FIXME: undeclared name: failOnError
	defer file.Close()
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: RegistrarJugada
// -> Escribe una nueva jugada para un jugador y una ronda dados
// Recibe: el jugador, la ronda y la jugada, todos como string
// Retorna: Nada
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
  
func registrarJugada(registro string){

	//TO-DO: Redefinir el input en caso de ser necesario
	//TO-DO: En caso de redefinir el input hacer de ronda y
	//jugador textos
	//NOTE: jugador_n__ronda_m.txt
	
	//Se abre el archivo
	s := strings.Fields(dato)
	var fileName string = s[0] + "__" + s[1] + ".txt"
	file, error1 := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//?: ¿Que es 0644?

	//FIXME: undeclared name: failOnError

	JugadaARegistar :=  s[2] + "\n"
	_, error2 := file.WriteString(JugadaARegistar)

	
	defer file.Close()
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// Función: DevolverJugadas
// -> Recibe un jugador y una ronda y devulve las jugadas
// presentes en el archivo correspondiente al input recibido
// Recibe: el jugador y la ronda todos como string
// Retorna: Un arreglo de enteros con las jugadas ordenadas
// en el mismo orden en que se registraron
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func DevolverJugadas(jugador string, ronda string) string{

	
	var jugadas string = ""
	var jugada string
	
	//TO-DO: Redefinir el input en caso de ser necesario
	//TO-DO: En caso de redefinir el input hacer de ronda y
	//jugador textos
	//NOTE: jugador_n__ronda_m.txt
	
	//Se abre el archivo
	var fileName string = jugador + "__" + ronda + ".txt"
	file, error1 := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	//FIXME: undeclared name: failOnError

	scanner := bufio.NewScanner(file)

	//Recorrer el archivo para registrar las jugadas una a una
    for scanner.Scan(){
		jugada = scanner.Text()
		jugadas = jugadas + jugada + " "
    }

	
	defer file.Close()
	return jugadas
}