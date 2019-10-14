package main

import (
	"log"
	"net"
	"context"
	"google.golang.org/grpc"


	"github.com/archacon/chacoworks/blockchain/proto"
	"github.com/archacon/chacoworks/blockchain/server/blockchain"
)
func main() {
	listener, err := net.Listen("tcp", "8080")
	if err != nil {
		log.Fatalf("unablet to listen on 8080 port :%v", err)

	}

	srv  := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &Server{
		Blockchain: blockchain.NewBlockchain
	})
	srv.Serve(listener)
}

type Server struct {
	BlockChain *blockchain.Blockchain
}

func (s *Server) AddBlock(ctx context.Context, in *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	block := s.Blockchain.AddBlock(in.Data)

	return new(proto.AddBlockResponse{
		Hash: block.Hash, 
	}),nil
}

func (s *Server) GetBlockchain(ctx context.Context, in *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	resp := new(proto.GetBlockchainResponse)
	for _, block := range s.Blockchain.Blocks {
		resp.Blocks = append(resp.Blocks, &proto.Block{
			PrevBlockHash: block.PrevBlockHash,
			Hash: block.Hash,
			Data: block.Data,
		})
	}
	return resp,nil
}