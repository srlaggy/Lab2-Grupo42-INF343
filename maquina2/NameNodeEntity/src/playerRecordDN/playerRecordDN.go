package playerRecordDN

import (
	"context"
	"fmt"
	pr "lab/namenode/proto/playerRecordDN"
	ut "lab/namenode/utils"
	"time"

	"google.golang.org/grpc"
)

const (
	protocolo_grpc = ""
	port_grpc = "60100"
)

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: enviar jugadas al datanode ----- // --> Namenode actua como cliente
func PlayerRecordNameNode(jugador string, game string, address string) string{
	// Creamos conexion
	conn, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, port_grpc), grpc.WithInsecure(), grpc.WithBlock())
	ut.FailOnError(err, "Failed to create a connection")
	defer conn.Close()

	// Creamos conexion con el servicio 
	cpr := pr.NewRecordServiceClient(conn)
	// Conectamos con el servidor y se imprime la respuesta
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp_record,err2 := cpr.SendRecord(ctx, &pr.RecordReq{Player: jugador, Game: game})
	ut.FailOnError(err2, "Failed to send a play")
	fmt.Printf("El registro del jugador es: %s \n", resp_record.GetRecord())
	return resp_record.GetRecord()

}