package main

import (
	"example-rpc/jsonrpc/service"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.RegisterName("HelloService", new(service.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
