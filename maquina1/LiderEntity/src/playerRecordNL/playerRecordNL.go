package playerRecordNL

import (
	"context"
	"time"
	// "fmt"
	"google.golang.org/grpc"
	ut "lab/lider/utils"
	pr "lab/lider/proto/playerRecordNL"
)

const (
	address = "10.6.43.46"
	protocolo_grpc = ""
	port_grpc = "60002"
)

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: enviar jugadas al NameNode ----- // --> Lider actua como cliente
// jugador: "jugador_3"
func PlayerRecordLider(jugador string) string{
	// Creamos conexion
	conn, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn.Close()

	// Creamos conexion con el servicio 
	cpr := pr.NewPlayerRecordServiceClient(conn)
	// Conectamos con el servidor y se imprime la respuesta
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp_record,err2 := cpr.PlayerRecord(ctx, &pr.PlayerReq{Player: jugador})
	ut.FailOnError(err2, "Failed to send a play")
	// fmt.Printf("El registro del jugador es: %s \n", resp_record.GetRecords())
	return resp_record.GetRecords()
}