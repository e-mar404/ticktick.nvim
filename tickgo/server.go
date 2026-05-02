package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/ugorji/go/codec"
)

var mh codec.MsgpackHandle

func startServer() {
	auth := new(Auth)
	rpc.Register(auth)
	
	// TODO get port from env var, but needs to be coordinated with the lua implementation
	log.Println("listening on :8080")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error listening: ", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatalln("Error accepting connection: ", err)
		}

		rpcCodec := codec.MsgpackSpecRpc.ServerCodec(conn, &mh)
		go rpc.ServeCodec(rpcCodec)
	}
}
