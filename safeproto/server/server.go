package main

import (
	"example-rpc/protorpc/service"
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

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error", err)
		}

		rpc.ServeConn(conn)
	}

}
