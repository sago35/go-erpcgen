package main

import (
	"fmt"
)

var (
	debug = true
)

func rpc_tcpip_adapter_init() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_init()\n")
	}
	msg := startWriteMessage(0x00, 0x0f, 0x01, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}
