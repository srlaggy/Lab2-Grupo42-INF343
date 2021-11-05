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
	// port_grpc = "60111"
	address = "localhost"
)

// --------------- FUNCIONES GRPC --------------- //

// ----- FUNCIÓN: enviar jugadas al datanode ----- // --> Namenode actua como cliente
func PlayerRecordNameNode(jugador string, game string, ip_datanode string) string{
	if (ip_datanode == "Datanode_1"){
		fmt.Println("He entrado a PlayerRecordGame con:", jugador, game)
		// Creamos conexion
		conn, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, "60111"), grpc.WithInsecure(), grpc.WithBlock())
		ut.FailOnError(err, "Failed to create a connection")
		defer conn.Close()
	
		// Creamos conexion con el servicio 
		cpr := pr.NewRecordServiceClient(conn)
		// Conectamos con el servidor y se imprime la respuesta
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		resp_record,err2 := cpr.SendRecord(ctx, &pr.RecordReq{Player: jugador, Game: game})
		ut.FailOnError(err2, "Failed to send a play NameNode")
		fmt.Printf("El registro del jugador es: %s \n", resp_record.GetRecord())
		return resp_record.GetRecord()
	}else if (ip_datanode == "Datanode_2"){
		fmt.Println("He entrado a PlayerRecordGame con:", jugador, game)
		// Creamos conexion
		conn, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, "60112"), grpc.WithInsecure(), grpc.WithBlock())
		ut.FailOnError(err, "Failed to create a connection")
		defer conn.Close()
	
		// Creamos conexion con el servicio 
		cpr := pr.NewRecordServiceClient(conn)
		// Conectamos con el servidor y se imprime la respuesta
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		resp_record,err2 := cpr.SendRecord(ctx, &pr.RecordReq{Player: jugador, Game: game})
		ut.FailOnError(err2, "Failed to send a play NameNode")
		fmt.Printf("El registro del jugador es: %s \n", resp_record.GetRecord())
		return resp_record.GetRecord()
	}else{
		fmt.Println("He entrado a PlayerRecordGame con:", jugador, game)
		// Creamos conexion
		conn, err := grpc.Dial(ut.CreateDir(protocolo_grpc, address, "60112"), grpc.WithInsecure(), grpc.WithBlock())
		ut.FailOnError(err, "Failed to create a connection")
		defer conn.Close()
	
		// Creamos conexion con el servicio 
		cpr := pr.NewRecordServiceClient(conn)
		// Conectamos con el servidor y se imprime la respuesta
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		resp_record,err2 := cpr.SendRecord(ctx, &pr.RecordReq{Player: jugador, Game: game})
		ut.FailOnError(err2, "Failed to send a play NameNode")
		fmt.Printf("El registro del jugador es: %s \n", resp_record.GetRecord())
		return resp_record.GetRecord()
	}


}