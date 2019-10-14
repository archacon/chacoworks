package main

import (
	"log"
	"net"
	"google.golang.org/grpc"


	"github.com/archacon/chacoworks/blockchain/proto"
)
func main() {
	listener, err := net.Listen("tcp", "8080")
	if err != nil {
		log.Fatalf("unablet to listen on 8080 port :%v", err)

	}

	srv  := grpc.NewServer()
}