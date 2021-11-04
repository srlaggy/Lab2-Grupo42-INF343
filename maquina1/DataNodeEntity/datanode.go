package main

import (
	dnr "lab/datanode/src/dataNodeRegistro"
	spd "lab/datanode/src/sendPlaysDN"
    prd "lab/datanode/src/playerRecordDN"
)

func main(){

	go spd.Grpc_func()
    prd.Grpc_func()
}