package main

import (
	"example-rpc/saferpc/service"
	"log"
	"net"
	"net/rpc"
)

func main() {
	service.RegisterHelloService(new(service.HelloService))

	listenser, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listenser.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go rpc.ServeConn(conn)
	}

}
