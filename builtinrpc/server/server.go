package main

import (
	"example-rpc/builtinrpc/service"
	"log"
	"net"
	"net/rpc"
)

func main() {
	rpc.RegisterName("HelloService", new(service.HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("accept error", err)
	}

	rpc.ServeConn(conn)

}
