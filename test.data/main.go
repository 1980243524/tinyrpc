package main

import (
	"log"
	"net"

	"test.data/message"

	"github.com/1980243524/tinyrpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}

	server := tinyrpc.NewServer()
	server.RegisterName("ArithService", new(message.ArithService))
	server.Serve(lis)
}
