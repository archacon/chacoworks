package main

import (
	"context"
	"flag"
	"log"
	"time"
	"github.com/archacon/chacoworks/blockchain/proto"
	"google.golang.org/grpc"
)

var client proto.BlockchainClient

func main() {
	addFlag := flag.Bool("add", false,"add new block")
	listFlag := flag.Bool("list", false, "get the blockchain")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}

	client = proto.NewBlockchainClient(conn)
	
	if *addFlag {
		addBlock()
	}
	
	if *listFlag {
		getBlockchain()
	}
}

func addBlock() {
	block, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})
	if err != nil {
		log.Fatalf("unable to add block: %v", err)
	}
	log.Printf("new block hash: %s\n", block.Hash)
}


func getBlockchain() {
	bc, err := client.GetBlockchain(context.Background(), &proto.GetBlockchainRequest{})
	if err != nil {
		log.Fatalf("unable to add blockchain: %v", err)
	}

	log.Println("blocks:\n")
	for _, b := range bc.Blocks {
		log.Printf("hash: %s, prev block hash: %s, data: %s", b.Hash, b.PrevBlockHash, b.Data)
	}
}