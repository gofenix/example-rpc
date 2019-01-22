package main

import (
	"example-rpc/saferpc/service"
	"fmt"
	"log"
)

func main() {
	client, err := service.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	var reply string
	err = client.Hello("你好啊，Lucas", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
