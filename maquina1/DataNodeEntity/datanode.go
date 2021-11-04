package main

import (
	prd "lab/datanode/src/playerRecordDN"
	spd "lab/datanode/src/sendPlaysDN"
)

func main(){

	go spd.Grpc_func()
    prd.Grpc_func()
}