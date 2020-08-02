package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/LuisAcerv/btchdwallet/proto/btchdwallet"
	"github.com/LuisAcerv/btchdwallet/wallet"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50055"
)

type server struct {
	pb.UnimplementedWalletServer
}

func (s *server) CreateWallet(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	fmt.Println()
	fmt.Println("\nCreating new wallet")

	wallet := wallet.CreateWallet()

	return wallet, nil
}

func (s *server) GetWallet(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	fmt.Println()
	fmt.Println("\nGetting wallet data")

	wallet := wallet.DecodeWallet(in.Mnemonic)

	return wallet, nil
}

func (s *server) GetBalance(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	fmt.Println()
	fmt.Println("\nGetting Balance data")

	balance := wallet.GetBalance(in.Address)

	return balance, nil
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
