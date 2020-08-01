package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/LuisAcerv/btchdwallet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50055"
)

type server struct {
	pb.UnimplementedWalletServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWalletServer(s, &server{})
	reflection.Register(s)
	fmt.Printf("Service running at port: %s", port)
	fmt.Println()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
