package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	RPC_PORT      string
	CALLBACK_PORT string
)

func main() {
	flag.StringVar(&RPC_PORT, "rpcPort", "", "port to start rpc server on")
	flag.StringVar(&CALLBACK_PORT, "callbackPort", "", "port to start callback server on")
	flag.Parse()

	if RPC_PORT == "" || CALLBACK_PORT == "" {
		fmt.Printf("Both flags rpcPort and callbackPort are required\n")
		os.Exit(1)
	}

	startRPCServer()
}
