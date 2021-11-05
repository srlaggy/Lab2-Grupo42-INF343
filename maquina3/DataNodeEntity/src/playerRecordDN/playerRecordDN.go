package playerRecordDN

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	rp "lab/datanode/proto/playerRecordDN"
	dr "lab/datanode/src/dataNodeRegistro"
	ut "lab/datanode/utils"
)

const (
	address = "localhost"
	protocolo_grpc = "tcp"
	port_grpc = "60112" //Cambiar puerto y direccion
)

// ----- FUNCIÓN: Devolver registro del jugador ----- // --> DataNode actua como servidor
type server struct {
	rp.UnimplementedRecordServiceServer
}

// funcion: conecta con el service Record
func (s *server) SendRecord(ctx context.Context, in *rp.RecordReq) (*rp.RecordResp, error) {
	log.Printf("Jugador y ronda %v", in.Player)
	return &rp.RecordResp{Record: dr.DevolverJugadas(in.Player, in.Game)}, nil
}

// funciones: crea la conexión
func Grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	rp.RegisterRecordServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v", port_grpc)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}