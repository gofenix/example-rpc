package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "this is json rpc demo", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)

}
