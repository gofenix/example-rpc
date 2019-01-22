package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dial error", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "你好，我是Lucas", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
