package playerRecordNL

import(
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	rp "lab/namenode/proto/playerRecordNL"
	ut "lab/namenode/utils"
)

const (
	address = "localhost"
	protocolo_grpc = "tcp"
	port_grpc = "60002"
)

// ----- FUNCIÓN: Devolver registro del jugador ----- // --> NameNode actua como servidor
type server struct {
	rp.UnimplementedPlayerRecordServiceServer
}

// funcion: conecta con el service PlayerRecord
func (s *server) PlayerRecord(ctx context.Context, in *rp.PlayerReq) (*rp.PlayerResp, error) {
	log.Printf("Received %v", in.Player) // 
	return &rp.PlayerResp{Records: "REGISTRO DE JUGADAS DEL JUGADOR"}, nil
}

// funciones: crea la conexión
func Grpc_func() {
	lis, err := net.Listen(protocolo_grpc, ":"+port_grpc)
	ut.FailOnError(err, "Failed to listen")

	s := grpc.NewServer()
	rp.RegisterPlayerRecordServiceServer(s, &server{})
	log.Printf("Servidor grpc escuchando en el puerto %v", port_grpc)
	ut.FailOnError(s.Serve(lis), "Failed to serve")
}